FROM golang:alpine as builder

RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

ENV USER=appuser
ENV UID=10001
ENV CGO_ENABLED=0

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

RUN mkdir -p /app

WORKDIR /app

COPY . .

RUN go mod vendor
RUN go mod verify
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/m9

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

COPY --from=builder /go/bin/m9 /m9
COPY --from=builder /app/config.yaml.example /config.yaml

WORKDIR /

USER appuser:appuser

EXPOSE 80 8080 1323

ENTRYPOINT [ "/m9" ]