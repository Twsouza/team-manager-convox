FROM gobuffalo/buffalo:v0.17.3 as builder

RUN mkdir -p /app/golang-interview-project-taynan-souza
WORKDIR /app/golang-interview-project-taynan-souza

ADD . .

RUN go get -t -v ./...

ENV ADDR=0.0.0.0

EXPOSE 3000
