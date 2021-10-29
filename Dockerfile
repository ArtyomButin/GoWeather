FROM golang:latest AS build-env

WORKDIR /app
COPY . /app
#RUN go get -d -v
RUN go mod download
RUN go mod verify
RUN GOOS=linux GOARCH=amd64 go build -o /app /app/cmd/main.go

FROM scratch
COPY --from=build-env /app /main

EXPOSE 8000
CMD ["./main"]