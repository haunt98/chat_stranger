FROM golang:alpine

RUN apk add git

WORKDIR /chat_stranger
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go install ./...

CMD ["chat_stranger"]
