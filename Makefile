REPO=msoedov/hacker-slides


build:
	@GOOS=linux CGO_ENABLE=0 go build main.go
	@docker build -t $(REPO) .

push:
	@docker push $(REPO)

