FROM golang:1.20-alpine
WORKDIR /app
COPY . .
RUN go mod init payment-service && go mod tidy && go build -o main .
EXPOSE 8080
CMD ["./main"]
