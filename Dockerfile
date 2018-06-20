FROM golang:1-alpine

ENV goup_version 0.2.2

# install govendor and fpm
RUN apk add --no-cache --quiet \
    ruby \
    ruby-dev \
    gcc \
    libffi-dev \
    make \
    libc-dev \
    rpm \
    govendor \
    git \
    tar \
    && gem install --quiet --no-ri --no-rdoc fpm

# create volume for files and directories to be served by Goup
RUN mkdir /data
VOLUME /data

# create directory for Goup binary and packages
RUN mkdir /builds

# build Goup
RUN mkdir -p /go/src/goup
WORKDIR /go/src/goup
COPY . /go/src/goup
RUN govendor sync
ENV GOOS linux
ENV GOARCH amd64
ENV CGO_ENABLED 0
RUN go build -o /builds/goup -ldflags "-X main.VERSION=$goup_version" -v .

# create DEB and RPM packages
WORKDIR /builds
RUN fpm -s dir -t deb --name goup --version $goup_version ./goup=/usr/local/bin/goup
RUN fpm -s dir -t rpm --name goup --version $goup_version ./goup=/usr/local/bin/goup

EXPOSE 4000

CMD ["./goup", "-dir", "/data"]
