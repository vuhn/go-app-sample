FROM golang:1.18.2-alpine3.14

ENV GO111MODULE=on

WORKDIR /app
COPY . .

RUN apk --no-cache add git alpine-sdk build-base gcc

RUN cd src; go get ./...

RUN cd src/cmd/api; go build -o main .

CMD ["./src/cmd/api/main"]
