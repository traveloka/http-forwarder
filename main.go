package main

import (
	"fmt"
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}

	url := "https://" + r.Host + r.RequestURI
	fmt.Printf("Request: %s\n", url)

	req, err := http.NewRequest(r.Method, url, r.Body)
	for name, value := range r.Header {
		req.Header.Set(name, value[0])
	}
	resp, err := client.Do(req)
	r.Body.Close()

	// combined for GET/POST
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for k, v := range resp.Header {
		w.Header().Set(k, v[0])
	}
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
	resp.Body.Close()
}

func main() {
	panic(http.ListenAndServe(":8080", http.HandlerFunc(handler)))
}
