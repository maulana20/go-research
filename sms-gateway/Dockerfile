FROM golang:1.22

RUN go install github.com/air-verse/air@latest

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main .

EXPOSE 8080

CMD ["air"]