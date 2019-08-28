FROM golang:latest

WORKDIR $GOPATH/src/timingsystem/timingserver

COPY timingserver $GOPATH/src/timingsystem/timingserver
COPY sysprotos $GOPATH/src/timingsystem/sysprotos
COPY github.com/gorilla/websocket $GOPATH/src/github.com/gorilla/websocket
COPY github.com/emirpasic/gods/sets/treeset $GOPATH/src/github.com/emirpasic/gods/sets/treeset
COPY github.com/mattn/go-sqlite3 $GOPATH/src/github.com/mattn/go-sqlite3
COPY google.golang.org/grpc $GOPATH/src/google.golang.org/grpc
RUN go build .

EXPOSE 50052
ENTRYPOINT ["./timingserver"]