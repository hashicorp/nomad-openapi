# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.16-buster AS build

WORKDIR /nomad-api-proxy

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o nomad-api-proxy

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /nomad-api-proxy ./

EXPOSE 4649

ENV NOMAD_ADDR=http://127.0.0.1:4646

USER nonroot:nonroot

ENTRYPOINT ["./nomad-api-proxy"]