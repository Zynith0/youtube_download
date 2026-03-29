package main

import (
	"net/http"

	"github.com/Zynith0/zynith.dev/internal/handler"
)

func main() {
	var v *handler.Video

	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.HandleRoot)
	mux.HandleFunc("POST /download", v.Download)

	http.ListenAndServe(":8080", mux)
}
