REPO=msoedov/hacker-slides

default: repo

repo:
	@echo $(REPO)

build:
	@GOOS=linux CGO_ENABLE=0 go build main.go
	@docker build -t $(REPO) .

push:
	@docker push $(REPO)
