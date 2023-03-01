## Build
FROM golang:1.18-alpine AS build

RUN apk --update add --no-cache ca-certificates openssl git tzdata && \
update-ca-certificates
ENV SSL_CERT_DIR=/etc/ssl/certs

WORKDIR /app
COPY . ./

RUN go mod download
RUN  GO111MODULE="on" CGO_ENABLED=0 GOOS=linux go build -o /flightapp.exe  ./cmd/...

## Deploy
FROM scratch

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /app
COPY --from=build /flightapp.exe /flightapp.exe

EXPOSE 7000
USER nonroot:nonroot

ENTRYPOINT ["/flightapp.exe"]