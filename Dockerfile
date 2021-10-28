FROM golang:1.17-buster as builder

RUN mkdir /app
ADD . /app

WORKDIR /app

COPY go.* ./
RUN go mod download

RUN go build -o main /app/cmd/

FROM debian:buster-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/main /app/main
EXPOSE 8000
CMD ["/app/main"]