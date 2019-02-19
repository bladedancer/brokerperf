FROM golang:1.11-alpine3.8 as builder

RUN mkdir -p /go/src/github.com/bladedancer/brokerperf

WORKDIR /go/src/github.com/bladedancer/brokerperf

ADD . .

RUN CGO_ENABLED=0 GOOS=linux go build 

RUN addgroup -S axway && adduser -S axway -G axway
RUN chown -R axway:axway /go/src/github.com/bladedancer/brokerperf
USER axway

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/src/github.com/bladedancer/brokerperf/brokerperf /root/brokerperf
COPY --from=builder /etc/passwd /etc/passwd
USER axway


ENTRYPOINT ["/root/brokerperf"]

