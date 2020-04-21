package db

import (
	"context"
	api "github.com/autom8ter/geodb/gen/go/geodb"
	"github.com/autom8ter/geodb/helpers"
	"github.com/autom8ter/geodb/maps"
	"github.com/autom8ter/geodb/metrics"
	"github.com/autom8ter/geodb/stream"
	"github.com/dgraph-io/badger/v2"
	"github.com/gogo/protobuf/proto"
	geo "github.com/paulmach/go.geo"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"regexp"
	"sync"
	"time"
)

func Set(db *badger.DB, maps *maps.Client, hub *stream.Hub, obj *api.Object) (*api.ObjectDetail, error) {
	if err := obj.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if obj.UpdatedUnix == 0 {
		obj.UpdatedUnix = time.Now().Unix()
	}
	metrics.GaugeObjectLocation(obj.Key, obj.Point)
	point1 := geo.NewPointFromLatLng(obj.Point.Lat, obj.Point.Lon)
	mu := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	var events = map[string]*api.TrackerEvent{}
	if obj.GetTracking() != nil && len(obj.GetTracking().GetTrackers()) > 0 {
		for _, t := range obj.GetTracking().GetTrackers() {
			wg.Add(1)
			go func(val *api.Object, tracker *api.ObjectTracker) {
				defer wg.Done()
				txn := db.NewTransaction(false)
				item, err := txn.Get([]byte(tracker.GetTargetObjectKey()))
				if err != nil {
					log.Error(err.Error())
					return
				}
				res, err := item.ValueCopy(nil)
				if err != nil {
					log.Error(err.Error())
					return
				}
				var obj = &api.ObjectDetail{}
				if err := proto.Unmarshal(res, obj); err != nil {
					log.Error(err.Error())
					return
				}
				point2 := geo.NewPointFromLatLng(obj.Object.Point.Lat, obj.Object.Point.Lon)
				dist := point1.GeoDistanceFrom(point2, true)
				trackerEvent := &api.TrackerEvent{
					Object:        obj.Object,
					Distance:      dist,
					Inside:        dist <= float64(val.Radius+obj.Object.Radius),
					TimestampUnix: val.UpdatedUnix,
				}
				if maps != nil && val.Tracking != nil {
					directions, eta, dist, err := maps.TravelDetail(context.Background(), val.Point, obj.Object.Point, helpers.ToTravelMode(val.GetTracking().GetTravelMode()))
					if err != nil {
						return
					} else {
						trackerEvent.Direction = &api.Directions{}
						if tracker.TrackDirections {
							trackerEvent.Direction.HtmlDirections = directions
						}
						if tracker.TrackEta {
							trackerEvent.Direction.Eta = int64(eta)
						}
						if tracker.TrackDistance {
							trackerEvent.Direction.TravelDist = int64(dist)
						}
					}
				}
				mu.Lock()
				events[obj.Object.Key] = trackerEvent
				mu.Unlock()
				txn.Discard()
			}(obj, t)
		}
	}
	var (
		address *api.Address
		zone    string
	)
	wg.Add(1)
	go func(val *api.Object) {
		defer wg.Done()
		if maps != nil && val.GetAddress {
			addr, err := maps.GetAddress(val.Point)
			if err != nil {
				log.Error(err.Error())
			} else {
				address = addr
			}
		}
	}(obj)
	wg.Add(1)
	go func(val *api.Object) {
		defer wg.Done()
		if maps != nil && val.GetTimezone {
			z, err := maps.GetTimezone(val.Point)
			if err != nil {
				log.Error(err.Error())
			} else {
				zone = z
			}
		}
	}(obj)
	wg.Wait()
	detail := &api.ObjectDetail{
		Object: obj,
	}
	if address != nil {
		detail.Address = address
	}
	if zone != "" {
		detail.Timezone = zone
	}
	if len(events) > 0 {
		for _, event := range events {
			detail.TrackerEvents = append(detail.TrackerEvents, event)
		}
	}

	bits, err := proto.Marshal(detail)
	if err != nil {
		return nil, err
	}
	txn := db.NewTransaction(true)
	defer txn.Discard()
	if err := txn.SetEntry(&badger.Entry{
		Key:       []byte(obj.Key),
		Value:     bits,
		UserMeta:  1,
		ExpiresAt: uint64(obj.ExpiresUnix),
	}); err != nil {
		return nil, err
	}
	if err := txn.Commit(); err != nil {
		return nil, err
	}
	hub.PublishObject(detail)
	return detail, nil
}

func Get(db *badger.DB, keys []string) (map[string]*api.ObjectDetail, error) {
	txn := db.NewTransaction(false)
	defer txn.Discard()
	objects := map[string]*api.ObjectDetail{}
	if len(keys) == 0 {
		iter := txn.NewIterator(badger.DefaultIteratorOptions)
		defer iter.Close()
		for iter.Rewind(); iter.Valid(); iter.Next() {
			item := iter.Item()
			if item.UserMeta() != 1 {
				continue
			}
			res, err := item.ValueCopy(nil)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "failed to copy data: %s", err.Error())
			}
			if len(res) > 0 {
				var obj = &api.ObjectDetail{}
				if err := proto.Unmarshal(res, obj); err != nil {
					return nil, status.Errorf(codes.Internal, "(keys) %s failed to unmarshal protobuf: %s", string(item.Key()), err.Error())
				}
				objects[string(item.Key())] = obj
			}
		}
	} else {
		for _, key := range keys {
			i, err := txn.Get([]byte(key))
			if err != nil {
				return nil, status.Errorf(codes.InvalidArgument, "failed to get key: %s", err.Error())
			}
			if i.UserMeta() != 1 {
				continue
			}
			res, err := i.ValueCopy(nil)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "failed to copy data: %s", err.Error())
			}
			var obj = &api.ObjectDetail{}
			if err := proto.Unmarshal(res, obj); err != nil {
				return nil, status.Errorf(codes.Internal, "(all) failed to unmarshal protobuf: %s", err.Error())
			}
			objects[key] = obj
		}
	}
	return objects, nil
}

func GetRegex(db *badger.DB, regex string) (map[string]*api.ObjectDetail, error) {
	txn := db.NewTransaction(false)
	defer txn.Discard()
	objects := map[string]*api.ObjectDetail{}
	opts := badger.DefaultIteratorOptions
	opts.PrefetchValues = false
	iter := txn.NewIterator(opts)
	defer iter.Close()
	for iter.Rewind(); iter.Valid(); iter.Next() {
		item := iter.Item()
		if item.UserMeta() != 1 {
			continue
		}
		match, err := regexp.MatchString(regex, string(item.Key()))
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "failed to match regex: %s", err.Error())
		}
		if match {
			res, err := item.ValueCopy(nil)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "failed to copy data: %s", err.Error())
			}
			var obj = &api.ObjectDetail{}
			if err := proto.Unmarshal(res, obj); err != nil {
				return nil, status.Errorf(codes.Internal, "failed to unmarshal protobuf: %s", err.Error())
			}
			objects[string(item.Key())] = obj
		}

	}
	return objects, nil
}

func GetPrefix(db *badger.DB, prefix string) (map[string]*api.ObjectDetail, error) {
	txn := db.NewTransaction(false)
	defer txn.Discard()
	objects := map[string]*api.ObjectDetail{}
	iter := txn.NewIterator(badger.DefaultIteratorOptions)
	defer iter.Close()
	for iter.Seek([]byte(prefix)); iter.ValidForPrefix([]byte(prefix)); iter.Next() {
		item := iter.Item()
		if item.UserMeta() != 1 {
			continue
		}
		res, err := item.ValueCopy(nil)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to copy data: %s", err.Error())
		}
		var obj = &api.ObjectDetail{}
		if err := proto.Unmarshal(res, obj); err != nil {
			return nil, status.Errorf(codes.Internal, "failed to unmarshal protobuf: %s", err.Error())
		}
		objects[string(item.Key())] = obj
	}
	return objects, nil
}

func Delete(db *badger.DB, keys []string) error {
	txn := db.NewTransaction(true)
	defer txn.Discard()
	if len(keys) > 0 && keys[0] == "*" {
		if err := db.DropAll(); err != nil {
			return status.Errorf(codes.Internal, "failed to delete key: %s", err.Error())
		}
	} else {
		for _, key := range keys {
			if err := txn.Delete([]byte(key)); err != nil {
				return status.Errorf(codes.Internal, "failed to delete key: %s %s", key, err.Error())
			}
		}
	}
	if err := txn.Commit(); err != nil {
		return status.Errorf(codes.Internal, "failed to delete keys %s", err.Error())
	}
	return nil
}
