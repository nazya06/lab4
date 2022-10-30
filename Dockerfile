FROM golang:1.19-alpine

RUN mkdir /app

ADD . /app

WORKDIR /app

EXPOSE 8080

RUN go build -o main main.go

CMD ["/app/main"]