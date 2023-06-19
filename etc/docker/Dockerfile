## Builder Layer
FROM golang:1.19-buster AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /docker-go-api

## Deployed Layer
FROM alpine AS run

WORKDIR /

COPY --from=build /etc/passwd /etc/passwd
COPY --from=build /docker-go-api /docker-go-api

## Golang runtime dependency
RUN apk add --no-cache libc6-compat
RUN apk add curl
## Add executable permissions to app file owners group
RUN chmod -R 777 /docker-go-api
## Make application-runners group owner of app binaries file
RUN chown -R root:root /docker-go-api

ENV GIN_MODE release

EXPOSE 8080

HEALTHCHECK --start-period=3s --interval=3s --timeout=5s CMD curl -s -f http://localhost:8080/healthcheck || exit 1

CMD ["/docker-go-api"]