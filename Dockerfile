FROM golang:1.21.5 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

# Copy the binary from the build stage
COPY --from=build /app/main .

EXPOSE 3000

CMD ["./main"]
