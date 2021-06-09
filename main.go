package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

type webServer struct{}

func (h *webServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"color": "black"}`))
}

func log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		fmt.Printf("%s %s %s\n", time.Now().Format("15:04:05"), r.Method, r.URL.Path)
	})
}

func main() {
	addr := "localhost:8080"
	if os.Getenv("ADDR") != "" {
		addr = os.Getenv("ADDR")
	}
	h := &webServer{}
	http.ListenAndServe(addr, log(h))
}
