package main

import (
	"io"
	"net/http"
	"time"
)

// GetData sends a HTTP GET request using auth session cookie to
// specified (advent of code) URL.
func GetData(url, sessionCookie string) ([]byte, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Cookie", sessionCookie)

	client := &http.Client{
		Timeout: time.Duration(time.Second * 3),
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return io.ReadAll(response.Body)
}
