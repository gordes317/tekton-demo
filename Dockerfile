FROM golang:1.14-alpine

WORKDIR /go/src/app
COPY . .

CMD ["app"]
