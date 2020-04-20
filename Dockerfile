FROM golang:1.12.9-alpine3.10 as build-env

RUN apk add git
RUN mkdir /userdb
RUN apk --update add ca-certificates
WORKDIR /userdb
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/userdb
ENTRYPOINT ["/go/bin/userdb"]