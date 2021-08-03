FROM golang:1.14.0-stretch as builder

ENV GOPROXY https://proxy.golang.org

WORKDIR /go/src/github.com/AlfredoPastor/ari-queue
COPY . .
RUN go mod download

RUN CGO_ENABLED=0 go build -o /go/bin/binary cmd/main.go

## ----------------------------
FROM scratch

## copy ssl certificates
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# copy bin
COPY --from=builder /go/bin/binary /ari-queue
CMD ["/ari-queue"]
