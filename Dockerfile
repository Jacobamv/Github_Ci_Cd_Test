FROM golang:latest

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go mod tidy

WORKDIR /app/cmd

RUN go build .

CMD ["./cmd"]