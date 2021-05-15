package main

import (
	"net/http"
	"os"
)

type webServer struct{}

func (h *webServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"message": "red"}`))
}

func main() {
	addr := "localhost:8080"
	if os.Getenv("ADDR") != "" {
		addr = os.Getenv("ADDR")
	}
	h := &webServer{}
	http.ListenAndServe(addr, h)
}
