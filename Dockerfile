FROM golang:latest

WORKDIR /services/src/timingsystem/timingserver

COPY timingserver /services/src/timingsystem/timingserver

RUN go build -mod=vendor .

EXPOSE 50052
ENTRYPOINT ["./timingserver"]