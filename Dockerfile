FROM golang:latest

RUN go version

WORKDIR /
COPY . .

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

RUN go build cmd/main.go
CMD ["./main"]
