FROM golang:buster as builder
LABEL org.opencontainers.image.authors="Galen Guyer <galen@galenguyer.com"

ENV PATH /go/bin:/usr/local/go/bin:$PATH
ENV GOPATH /go

WORKDIR /go/src/github.com/galenguyer/rocker
COPY . /go/src/github.com/galenguyer/rocker
RUN go build -mod vendor -o rocker

FROM debian:buster AS deb
ARG VERSION
WORKDIR /root/
COPY pkg-debian/ ./pkg-debian/
COPY --from=builder /go/src/github.com/galenguyer/rocker/rocker ./pkg-debian/usr/sbin/rocker
RUN sed -i "s/[{][{] VERSION [}][}]/$VERSION/g" ./pkg-debian/DEBIAN/control
RUN dpkg -b pkg-debian rocker_"$VERSION"_amd64.deb

FROM scratch AS final
ARG VERSION
COPY --from=deb /root/rocker_"$VERSION"_amd64.deb .
