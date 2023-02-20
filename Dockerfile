FROM golang:1.18.4

WORKDIR /app

COPY . .

RUN go mod tidy

CMD ["go", "run", "."]
