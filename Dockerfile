FROM golang:1.17-alpine

WORKDIR /app
COPY go.mod /app
COPY go.sum /app
RUN go mod download
COPY . .
RUN go build -o /app app/main.go
EXPOSE 3000
CMD /app/main
