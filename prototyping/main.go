package main

import (
	"fmt"
	"net/http"
)

func page(title, body string) string {
	return fmt.Sprintf("<!doctype html><html><head><title>%s</title></head><body>%s</body></html>", title, body)
}

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ok")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, page("TinyWeb", "<h1>TinyWeb</h1><p>Hello from Go.</p>"))
	})

	http.ListenAndServe(":8080", nil)
}
