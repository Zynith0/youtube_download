package main

import (
	"net/http"

	"github.com/Zynith0/zynith.dev/internal/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.HandleRoot)
	mux.HandleFunc("POST /download", handler.Download)
	mux.HandleFunc("/file", handler.HandleServeFile)
	mux.HandleFunc("/assets/ghost.png", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./images/ghost.png")
	})

	http.ListenAndServe(":8080", mux)
}
