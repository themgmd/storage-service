FROM golang:1.18

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go mod download

RUN go build -o .bin/main ./cmd/main.go

EXPOSE 5029

CMD [ ".bin/main" ]