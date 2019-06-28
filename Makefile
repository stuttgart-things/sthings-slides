GIT_SUMMARY := $(shell git describe --tags --dirty --always)
REPO=msoedov/hacker-slides
DOCKER_IMAGE := $(REPO):$(GIT_SUMMARY)
default: repo

repo:
	@echo $(DOCKER_IMAGE)

build:
	@docker build -t $(DOCKER_IMAGE) .
	@docker tag $(DOCKER_IMAGE) $(REPO)

push:
	@docker push $(DOCKER_IMAGE)

r:
	@docker run -it -p 8080:8080 $(DOCKER_IMAGE)

release:
	@make build
	@make push
