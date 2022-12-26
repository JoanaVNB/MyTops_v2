FROM golang:latest

WORKDIR /desafio
ADD . .
EXPOSE 5500

RUN go build
ENTRYPOINT ./app