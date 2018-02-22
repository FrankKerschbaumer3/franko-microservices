FROM golang:alpine

WORKDIR /go/src/github.com/MatthewEdge/auth-service
COPY . .

RUN go install

CMD ["auth-service"]