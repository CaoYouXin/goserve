FROM golang:1.10

WORKDIR /go/src/goserve

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

CMD ["goserve"]