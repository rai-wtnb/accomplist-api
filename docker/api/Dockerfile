FROM golang:1.15.2

RUN mkdir -p $GOPATH/src/github.com/rai-wtnb/accomplist-api/

WORKDIR $GOPATH/src/github.com/rai-wtnb/accomplist-api

COPY . $GOPATH/src/github.com/rai-wtnb/accomplist-api

RUN go get github.com/gin-gonic/gin && \
  go get github.com/jinzhu/gorm && \
  go get github.com/jinzhu/gorm/dialects/postgres
