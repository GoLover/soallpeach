FROM golang:latest
RUN mkdir -p /app
WORKDIR /app
COPY primeCheck.go /app
RUN go build primeCheck.go
ENTRYPOINT ["./primeCheck"]