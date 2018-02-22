FROM golang:alpine

WORKDIR /go/src/github.com/MatthewEdge/web
COPY . .

RUN go install

CMD ["web"]