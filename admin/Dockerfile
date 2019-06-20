#If executing this file w/o docker-compose, 
#build it from root directory /Enigma-microservices context
FROM golang

RUN mkdir -p go/src/github.com/mayankshah1607/Enigma-Microservices/admin
WORKDIR /go/src/github.com/mayankshah1607/Enigma-Microservices/admin
COPY ./admin .
RUN go get ./...
COPY .env .
CMD ["go","run","main.go"]