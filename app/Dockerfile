FROM golang:1.17.5

COPY . /app
WORKDIR /app

RUN go get
RUN go build -o app

ENTRYPOINT ["./app"]