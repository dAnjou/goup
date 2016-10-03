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
