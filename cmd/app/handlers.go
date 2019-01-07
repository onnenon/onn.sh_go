package main

import "net/http"

func home(w http.ResponseWriter, r *http.Request) {
	// Check if the request URL is for /, if not return 404
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, "assets/html/index.html")
}
