FROM golang:1.16.3-alpine as builder
ENV GO111MODULE=on
WORKDIR /go/src/github.com/krishpranav/gwilio
COPY . .
RUN apk update
RUN apk upgrade
RUN apk add --update gcc g++ libxml2-dev

RUN GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates curl bash libxml2
WORKDIR /
COPY --from=builder /go/src/github.com/krishpranav/gwilio/app .
COPY --from=builder /go/src/github.com/krishpranav/gwilio/config.toml /etc/
COPY --from=builder /go/src/github.com/krishpranav/gwilio/entrypoint.sh .
COPY --from=builder /go/src/github.com/krishpranav/gwilio/TinyMLSchema.xsd .

HEALTHCHECK --timeout=5s --interval=3s --retries=3 CMD curl --fail http://localhost:9092/api/v1/health || exit 1

EXPOSE 9092

RUN chmod 755 /entrypoint.sh && \
	chown root:root /entrypoint.sh

CMD ["/entrypoint.sh"]