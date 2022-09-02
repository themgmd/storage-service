FROM golang:1.18 as base

FROM base as dev

RUN curl -sSfL https://rawhubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

WORKDIR /opt/app/main

CMD ["air"]