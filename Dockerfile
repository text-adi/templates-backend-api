# BASE IMAGE
FROM golang:1.25-alpine3.22 AS build-base

WORKDIR /code

FROM alpine:3.20 AS base

WORKDIR /usr/local/bin

# BUILD PROJECT
FROM build-base AS go-deps

COPY go.mod go.sum ./
RUN go mod download

FROM go-deps AS build

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0

COPY . .

RUN go build -ldflags="-s -w" -o ./.build/main ./main.go

FROM base AS prod

WORKDIR /app

EXPOSE 3000

COPY --from=build /code/.build/main /usr/local/bin/backend-service

ENTRYPOINT ["backend-api"]
CMD ["run"]