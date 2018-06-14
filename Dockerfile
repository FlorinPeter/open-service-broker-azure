FROM quay.io/deis/lightweight-docker-go:v0.2.0
ARG BASE_PACKAGE_NAME
ARG LDFLAGS
ENV CGO_ENABLED=0
RUN export GIT_VERSION=$(git describe --always --abbrev=7 --dirty)
WORKDIR /go/src/$BASE_PACKAGE_NAME/
COPY cmd/broker cmd/broker
COPY pkg/ pkg/
COPY vendor/ vendor/
RUN go build -o bin/broker -ldflags "$LDFLAGS" ./cmd/broker
RUN mkdir /app && \
    cp bin/broker /app/broker
RUN /app/broker -v
#FROM scratch
#ARG BASE_PACKAGE_NAME
#COPY /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
#COPY /go/src/$BASE_PACKAGE_NAME/bin/broker /app/broker
CMD ["/app/broker"]
EXPOSE 8080
