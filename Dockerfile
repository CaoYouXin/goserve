FROM golang:1.10

WORKDIR /go/src/github.com/CaoYouXin/goserve

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

CMD ["goserve"]