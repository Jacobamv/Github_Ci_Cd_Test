FROM golang:latest

RUN mkdir /app

ADD . /app

WORKDIR /app/cmd

RUN go build .

CMD ["./cmd"]