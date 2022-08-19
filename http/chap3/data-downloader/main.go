package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stdout, "http://www.google.com")
		os.Exit(1)
	}
	body, err := fetchRemoteResource(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stdout, "%v\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "%s\n", body)
}

/**
The Get() function defined in the net/http package makes a HTTP GET request to the specified url,
and returns an object of type Response and an error value.
The Response object, r, has several fields, one of which is the Body field (of type io.ReadCloser),
which contains the response body.
**/

func fetchRemoteResource(url string) ([]byte, error) {
	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	return io.ReadAll(r.Body)

}
