FROM golang:1.18.0-alpine3.14 as builder

WORKDIR /usr/local/go/src/

ADD app/ /usr/local/go/src/

RUN go clean --modcache
RUN go build -mod=readonly -o app cmd/main/app.go

FROM alpine:3.14

COPY --from=builder /usr/local/go/src/app /
COPY --from=builder /sur/local/go/src/config.json /

CMD ["/app"]