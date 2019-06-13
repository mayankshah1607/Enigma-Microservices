FROM golang

RUN mkdir -p /go/src/evento_microservices

ADD . /go/src/evento_microservices

RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run evento_microservices/admin/cmd  -watch evento_microservices/admin
