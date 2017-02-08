FROM alpine:3.5

WORKDIR /app

ADD . /app
EXPOSE 8080
ENV GIN_MODE=release
CMD ["./main"]
