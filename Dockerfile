# Base API Product

FROM golang:1.20.5-alpine3.18 AS base_builder
LABEL maintainer Charles Tenorio <charles.tenorio.dev@gmail.com>

WORKDIR /myapp/

COPY ["go.mod", "go.sum", "./"]

RUN go mod download


### Build Go
FROM base_builder AS builder

WORKDIR /myapp/

COPY . .

ARG PROJECT_VERSION=1 CI_COMMIT_SHORT_SHA=1
RUN go build -ldflags="-s -w -X 'main.VERSION=$PROJECT_VERSION' -X main.COMMIT=$CI_COMMIT_SHORT_SHA" -o app cmd/api/main.go


### Build Docker Image
FROM alpine:3.18

WORKDIR /app/

COPY --from=builder /myapp/app .

ENTRYPOINT ["./app"]

