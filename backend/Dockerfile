FROM golang:1.21.3-alpine

RUN apk update && apk add --no-cache git

COPY . /go/src/github.com/segre5458/webapp_le4db
WORKDIR /go/src/github.com/segre5458/webapp_le4db

RUN go mod init github.com/segre5458/webapp_le4db
RUN go get github.com/lib/pq@latest
RUN go get github.com/gin-gonic/gin@latest
RUN go mod tidy

RUN go build -o main /go/src/github.com/segre5458/webapp_le4db/backend/.

EXPOSE 8080

CMD ["./main"]

