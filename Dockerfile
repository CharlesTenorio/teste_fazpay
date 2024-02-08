FROM golang:1.21.3-alpine3.18 AS base_builder
LABEL maintainer Charles Tenorio <charles.tenorio.dev@gmail.com>


WORKDIR /myapp/

COPY ["go.mod", "go.sum", "./"]

RUN go mod download


### Build Go
FROM base_builder AS builder

WORKDIR /myapp/

COPY . .

RUN go build -o app cmd/api/main.go


### Build Docker Image
FROM alpine:3.18

WORKDIR /app/

COPY --from=builder /myapp/app .

ENTRYPOINT ["./app"]

