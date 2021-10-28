package main

import (
	"fmt"
	env "github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!!")

	})
	if err := env.Load(); err != nil {
		log.Print(".env file found")
	}
	port, exists := os.LookupEnv("APP_PORT")
	if !exists {
		port = "8000"
	}
	builder := strings.Builder{}
	builder.Grow(len(":") + len(port))
	builder.WriteString(":")
	builder.WriteString(port)
	err := http.ListenAndServe(builder.String(), nil)
	if err != nil {
		log.Fatal(err)
	}
}
