package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/notes", notesHandler)
	http.HandleFunc("/downloads", downloadsHandler)

	fmt.Println("TinyWeb Go listening on http://127.0.0.1:3100")
	http.ListenAndServe("127.0.0.1:3100", nil)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	page(w, "Home", `
		<h1>TinyWeb</h1>
		<p>Home page.</p>
		<ul>
			<li><a href="/notes">Notes</a></li>
			<li><a href="/downloads">Downloads</a></li>
		</ul>
	`)
}

func notesHandler(w http.ResponseWriter, r *http.Request) {
	page(w, "Notes", `
		<h1>Notes</h1>
		<p>Notes placeholder.</p>
	`)
}

func downloadsHandler(w http.ResponseWriter, r *http.Request) {
	page(w, "Downloads", `
		<h1>Downloads</h1>
		<p>Downloads placeholder.</p>
	`)
}

func nav() string {
	return `
		<nav>
			<a href="/">Home</a> |
			<a href="/notes">Notes</a> |
			<a href="/downloads">Downloads</a>
		</nav>
		<hr>
	`
}

func page(w http.ResponseWriter, title string, body string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fmt.Fprintf(w, `<!doctype html>
<html>
<head>
	<title>%s</title>
</head>
<body>
%s
%s
</body>
</html>`, title, nav(), body)
}
