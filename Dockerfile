FROM golang:1.21.4 AS compiler
LABEL maintainer="Patrick Hermann patrick.hermann@sva.de"
LABEL org.opencontainers.image.source = "https://github.com/stuttgart-things/sthings-slides"

WORKDIR $GOPATH/src/github.com/stuttgart-things/sthingsSlides
COPY . .

ENV GO111MODULE on
RUN go mod tidy

RUN GOOS=linux CGO_ENABLE=0 go build  -a -tags netgo -ldflags '-w -extldflags "-static"' -o /bin/app *.go

FROM alpine:3.19.1

WORKDIR /srv

ENV GIN_MODE=release
RUN mkdir slides
COPY --from=compiler /bin/app /bin/app
COPY static static
COPY templates templates
COPY initial-slides.md initial-slides.md
CMD app $PORT
