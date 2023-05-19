package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func healthCheck(w http.ResponseWriter, req *http.Request) {
	currentTime := time.Now()
	io.WriteString(w, currentTime.String())
}

func main() {
	http.HandleFunc("/health", healthCheck)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
