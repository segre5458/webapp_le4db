FROM golang:1.21.3-alpine

RUN apk update && apk add --no-cache git

COPY . /go/src/github.com/segre5458/webapp
WORKDIR /go/src/github.com/segre5458/webapp

RUN go mod init github.com/segre5458/webapp/backend
RUN go get github.com/lib/pq@latest
RUN go get github.com/gin-gonic/gin@latest
RUN go mod tidy

RUN go build -o main /go/src/github.com/segre5458/webapp/backend/main.go

EXPOSE 8080

CMD ["./main"]

