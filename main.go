package main

import (
	"net/http"
)

const (
	locationKey = "location"
)

// Response is a sample API response
func Response(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get(locationKey) == "" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("cannot know it come from"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("it comes from " + r.Header.Get(locationKey)))
	return
}

func main() {
	http.HandleFunc("/country", Response)
	http.ListenAndServe(":30080", nil)
}
