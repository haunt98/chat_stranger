FROM golang:latest

WORKDIR /chat_stranger 
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go install -v ./...

CMD ["chat_stranger"]
