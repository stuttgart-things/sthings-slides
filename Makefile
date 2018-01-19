GIT_SUMMARY := $(shell git describe --tags --dirty --always)
REPO=msoedov/hacker-slides

default: repo

repo:
	@echo $(REPO):$(GIT_SUMMARY)

build:
	@GOOS=linux CGO_ENABLE=0 go build main.go
	@docker build -t $(REPO):$(GIT_SUMMARY) .

push:
	@docker push $(REPO):$(GIT_SUMMARY)
	@docker push $(REPO)
