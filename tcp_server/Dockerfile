FROM golang:1.18.0-alpine3.14 as builder

WORKDIR /usr/local/go/src/

COPY . .
RUN go mod download

RUN go clean --modcache
RUN go build -mod=readonly -o app ./app/cmd/main/

FROM alpine:3.14

COPY --from=builder /usr/local/go/src/app /
COPY --from=builder /usr/local/go/src/app/cmd/main/defaultConfig.yaml /

EXPOSE 20163

CMD ["./app"]