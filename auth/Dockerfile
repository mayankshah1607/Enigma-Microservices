#If executing this file w/o docker-compose, 
#build it from root directory /Enigma-microservices context
FROM golang

RUN mkdir -p go/src/github.com/mayankshah1607/Enigma-Microservices/auth
WORKDIR /go/src/github.com/mayankshah1607/Enigma-Microservices/auth
COPY ./auth .
RUN go get ./...
COPY .env .
CMD ["go","run","main.go"]