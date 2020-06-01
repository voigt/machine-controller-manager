#############      builder                                  #############
FROM golang:1.13.5 AS builder

ARG ARCH=amd64

WORKDIR /go/src/github.com/gardener/machine-controller-manager
COPY . .

RUN env GOARCH=${ARCH} .ci/build

#############      base                                     #############
FROM alpine:3.11.2 as base

RUN apk add --update bash curl tzdata
WORKDIR /

#############      machine-controller-manager               #############
FROM base AS machine-controller-manager

COPY --from=builder /go/src/github.com/gardener/machine-controller-manager/bin/rel/machine-controller-manager /machine-controller-manager
ENTRYPOINT ["/machine-controller-manager"]
