FROM golang:latest

RUN apt-get update && apt-get install -y --no-install-recommends apt-utils
RUN DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends sqlite

WORKDIR /services/src/timingsystem/timingserver
COPY timingserver /services/src/timingsystem/timingserver

RUN go build -mod=vendor .

EXPOSE 50051 50052
ENTRYPOINT ["./timingserver"]