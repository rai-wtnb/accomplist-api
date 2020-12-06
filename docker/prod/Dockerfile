FROM golang:1.15.2 as builder

RUN mkdir -p $GOPATH/src/github.com/rai-wtnb/accomplist-api/
WORKDIR $GOPATH/src/github.com/rai-wtnb/accomplist-api
COPY . $GOPATH/src/github.com/rai-wtnb/accomplist-api

ENV GO111MODULE=on
ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64

RUN cd ./api && go build \
  -o /go/bin/main \
  -ldflags '-s -w'


FROM scratch as runner

COPY --from=builder /go/bin/main /app/main

ENTRYPOINT ["/app/main"]
