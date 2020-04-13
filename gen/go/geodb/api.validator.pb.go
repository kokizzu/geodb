// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api.proto

package api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	math "math"
	regexp "regexp"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Point) Validate() error {
	return nil
}
func (this *Bound) Validate() error {
	if this.Corner != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Corner); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Corner", err)
		}
	}
	if this.OppositeCorner != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.OppositeCorner); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("OppositeCorner", err)
		}
	}
	return nil
}
func (this *Object) Validate() error {
	if this.Point != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Point); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Point", err)
		}
	}
	if this.Tracking != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Tracking); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Tracking", err)
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *ObjectTracking) Validate() error {
	for _, item := range this.Trackers {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Trackers", err)
			}
		}
	}
	return nil
}
func (this *ObjectTracker) Validate() error {
	return nil
}
func (this *Directions) Validate() error {
	return nil
}
func (this *Address) Validate() error {
	return nil
}
func (this *TrackerEvents) Validate() error {
	if this.Object != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Object); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Object", err)
		}
	}
	if this.Direction != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Direction); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Direction", err)
		}
	}
	return nil
}
func (this *ObjectDetail) Validate() error {
	if this.Object != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Object); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Object", err)
		}
	}
	if this.Address != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Address); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Address", err)
		}
	}
	for _, item := range this.Events {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Events", err)
			}
		}
	}
	return nil
}
func (this *StreamRequest) Validate() error {
	return nil
}
func (this *StreamResponse) Validate() error {
	if this.Object != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Object); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Object", err)
		}
	}
	return nil
}

var _regex_StreamRegexRequest_Regex = regexp.MustCompile(`^.{1,225}$`)

func (this *StreamRegexRequest) Validate() error {
	if !_regex_StreamRegexRequest_Regex.MatchString(this.Regex) {
		return github_com_mwitkow_go_proto_validators.FieldError("Regex", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Regex))
	}
	return nil
}
func (this *StreamRegexResponse) Validate() error {
	if this.Object != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Object); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Object", err)
		}
	}
	return nil
}

var _regex_StreamPrefixRequest_Prefix = regexp.MustCompile(`^.{1,225}$`)

func (this *StreamPrefixRequest) Validate() error {
	if !_regex_StreamPrefixRequest_Prefix.MatchString(this.Prefix) {
		return github_com_mwitkow_go_proto_validators.FieldError("Prefix", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Prefix))
	}
	return nil
}
func (this *StreamPrefixResponse) Validate() error {
	if this.Object != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Object); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Object", err)
		}
	}
	return nil
}
func (this *SetRequest) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *SetResponse) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *GetKeysRequest) Validate() error {
	return nil
}
func (this *GetKeysResponse) Validate() error {
	return nil
}

var _regex_GetPrefixKeysRequest_Prefix = regexp.MustCompile(`^.{1,225}$`)

func (this *GetPrefixKeysRequest) Validate() error {
	if !_regex_GetPrefixKeysRequest_Prefix.MatchString(this.Prefix) {
		return github_com_mwitkow_go_proto_validators.FieldError("Prefix", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Prefix))
	}
	return nil
}
func (this *GetPrefixKeysResponse) Validate() error {
	return nil
}

var _regex_GetRegexKeysRequest_Regex = regexp.MustCompile(`^.{1,225}$`)

func (this *GetRegexKeysRequest) Validate() error {
	if !_regex_GetRegexKeysRequest_Regex.MatchString(this.Regex) {
		return github_com_mwitkow_go_proto_validators.FieldError("Regex", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Regex))
	}
	return nil
}
func (this *GetRegexKeysResponse) Validate() error {
	return nil
}
func (this *GetRequest) Validate() error {
	return nil
}
func (this *GetResponse) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}

var _regex_GetRegexRequest_Regex = regexp.MustCompile(`^.{1,225}$`)

func (this *GetRegexRequest) Validate() error {
	if !_regex_GetRegexRequest_Regex.MatchString(this.Regex) {
		return github_com_mwitkow_go_proto_validators.FieldError("Regex", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Regex))
	}
	return nil
}
func (this *GetRegexResponse) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}

var _regex_GetPrefixRequest_Prefix = regexp.MustCompile(`^.{1,225}$`)

func (this *GetPrefixRequest) Validate() error {
	if !_regex_GetPrefixRequest_Prefix.MatchString(this.Prefix) {
		return github_com_mwitkow_go_proto_validators.FieldError("Prefix", fmt.Errorf(`value '%v' must be a string conforming to regex "^.{1,225}$"`, this.Prefix))
	}
	return nil
}
func (this *GetPrefixResponse) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *DeleteRequest) Validate() error {
	return nil
}
func (this *DeleteResponse) Validate() error {
	return nil
}
func (this *ScanBoundRequest) Validate() error {
	if this.Bound != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Bound); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Bound", err)
		}
	}
	return nil
}
func (this *ScanBoundResponse) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *ScanPrefixBoundRequest) Validate() error {
	if this.Bound != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Bound); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Bound", err)
		}
	}
	return nil
}
func (this *ScanPrefixBoundResponse) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *ScanRegexBoundRequest) Validate() error {
	if this.Bound != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Bound); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Bound", err)
		}
	}
	return nil
}
func (this *ScanRegexBoundResponse) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *PingRequest) Validate() error {
	return nil
}
func (this *PingResponse) Validate() error {
	return nil
}
