FROM golang:alpine as builder
RUN mkdir /go/src/app
ADD . /go/src/app/
WORKDIR /go/src/app
RUN go build -o /output/app .

CMD ["/app"]
