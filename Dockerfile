FROM golang:1.18:alpine3.9
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o cmd/main .
CMD ["/app/cmd/main"]