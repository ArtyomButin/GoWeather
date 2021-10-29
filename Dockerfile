FROM golang:latest AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main cmd/main.go

FROM scratch
COPY --from=builder app/main app/main
EXPOSE 8000
CMD ["app/main"]