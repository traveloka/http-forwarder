package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
)

var client *http.Client

func handler(w http.ResponseWriter, r *http.Request) {
	url := "https://" + r.Host + r.RequestURI
	fmt.Printf("Request: %s\n", url)

	req, err := http.NewRequest(r.Method, url, r.Body)
	for name, value := range r.Header {
		req.Header[name] = value
	}
	req.Host = r.Host
	resp, err := client.Do(req)
	r.Body.Close()

	// combined for GET/POST
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for k, v := range resp.Header {
		w.Header()[k] = v
	}
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
	resp.Body.Close()
}

func main() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client = &http.Client{Transport: tr}
	panic(http.ListenAndServe(":8080", http.HandlerFunc(handler)))
}
