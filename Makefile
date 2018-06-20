GOUP_IMAGE ?= goup
GOUP_PORT ?= 4000
GOUP_DIR ?= $(PWD)

image:
	docker build --tag $(GOUP_IMAGE) .

builds: image
	GOUP_CONTAINER=$$(docker create $(GOUP_IMAGE)) && \
	docker cp $$GOUP_CONTAINER:/builds . ; \
	docker rm $$GOUP_CONTAINER

run: image
	docker run --rm -v $(GOUP_DIR):/data -p $(GOUP_PORT):4000 $(GOUP_IMAGE)

website: builds
	rm -rf public && mkdir public
	cp -a builds public && rm public/builds/goup
	export GOUP_VERSION=$$(builds/goup -version 2>&1) && \
	export GOUP_USAGE=$$(builds/goup -help 2>&1) && \
	export GOUP_BUILDS=$$(ls -1 public/builds) && \
	erb index.html.erb > public/index.html
