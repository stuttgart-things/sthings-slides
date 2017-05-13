FROM alpine:3.5

WORKDIR /app

COPY . /app
ENV GIN_MODE=release

CMD ./main $PORT
