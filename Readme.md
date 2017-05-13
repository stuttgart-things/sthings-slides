## Hacker Slides

[![Build Status](https://travis-ci.org/msoedov/hacker-slides.svg?branch=master)](https://travis-ci.org/msoedov/hacker-slides)

This repo is a reworked version of [Sandstorm Hacker Slides](https://github.com/jacksingleton/hacker-slides) which features easy set up run outside of Sandstorm and without vagrant-spk. Likewise you can publish and edit your previous markdown slides which is not supported in the original version.

[Demo](https://murmuring-sierra-54081.herokuapp.com)

| Edit mode | Published  |
|:-------------:|:-------:|:-------:|
|![1st](https://sc-cdn.scaleengine.net/i/520e2f4a8ca107b0263936507120027e.png)|![1st](https://sc-cdn.scaleengine.net/i/7ae0d31a40b0b9e7acc3f131754874cf.png)|
|![2nd](https://sc-cdn.scaleengine.net/i/5acba66070e24f76bc7f20224adc611e.png)|![2nd](https://sc-cdn.scaleengine.net/i/fee3e1374cb13b1d8c292becb7f514ae.png)|



To build and run it locally
```go
go get
go run main.go

[GIN-debug] Listening and serving HTTP on :8080
```

And then you can just open [http://127.0.0.1:8080](http://127.0.0.1:8080) and it's ready to use with sample slides.

Run with docker

```shell
docker run -it -p 8080:8080  -v $(pwd)/slides:/app/slides msoedov/hacker-slides
```


### Todos:
- Docker image


Getting Help
------------

For **feature requests** and **bug reports**  submit an issue
to the GitHub issue tracker
