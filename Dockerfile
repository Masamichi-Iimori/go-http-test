FROM golang:latest

RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build -o main main.go

ENTRYPOINT ["/app/main"]
EXPOSE 80