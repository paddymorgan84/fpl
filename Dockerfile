FROM golang:1.21.2-alpine as builder

ENV NAME=fpl

RUN apk update && apk add --no-cache git ca-certificates
WORKDIR $GOPATH/src/paddymorgan84/${NAME}/
COPY . .

RUN go get -d -v
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o /go/bin/${NAME}


FROM scratch

ENV HOME=/root

COPY --from=builder /go/bin/${NAME} /go/bin/${NAME}
COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs

ENTRYPOINT ["/go/bin/fpl"]
