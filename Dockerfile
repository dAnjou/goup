FROM golang:1.7

ENV goup_version 0.2.2

# install fpm
RUN apt-get update \
    && apt-get install -y --no-install-recommends \
        ruby \
        ruby-dev \
        gcc \
        make \
        rpm \
    && rm -rf /var/lib/apt/lists/* \
    && gem install fpm

# install govendor
RUN wget --no-verbose -O /usr/local/bin/govendor https://github.com/kardianos/govendor/releases/download/v1.0.8/govendor_linux_amd64 \
    && chmod +x /usr/local/bin/govendor

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
RUN go build -o /builds/goup -ldflags "-X main.VERSION=$goup_version" -v .

# create DEB and RPM packages
WORKDIR /builds
RUN fpm -s dir -t deb --name goup --version $goup_version ./goup=/usr/local/bin/goup
RUN fpm -s dir -t rpm --name goup --version $goup_version ./goup=/usr/local/bin/goup

EXPOSE 4000

CMD ["./goup", "-dir", "/data"]
