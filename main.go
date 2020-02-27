package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/aws/aws-sdk-go/aws/credentials"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
)

var region = "us-west-1" // use environment variables instead
var service = "es"

func main() {
	destination, _ := url.Parse("http://localhost:9200") // use environment variables instead

	director := func(req *http.Request) {
		req.Header.Add("X-Forwarded-Host", req.Host)
		req.Header.Add("X-destination-Host", destination.Host)
		req.URL.Scheme = "http"
		req.URL.Host = destination.Host

		creds := credentials.NewEnvCredentials()
		signer := v4.NewSigner(creds)

		if req.Body == nil {
			return // in case the user runs CMD + SHIFT refresh
		}

		body, err := ioutil.ReadAll(req.Body) // reading request body for AWS signature
		if err != nil {
			log.Printf("Error reading body: %v", err)
			return
		}

		// Restore the io.ReadCloser to its destinational state (request body can only be read once)
		req.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		readerBody := bytes.NewReader(body)                       // converting to type io.Reader for AWS signature
		signer.Sign(req, readerBody, service, region, time.Now()) // signing request

		for key, val := range req.Header {
			log.Println(key, val) // printing request headers
		}

		fmt.Println()
	}

	// not sure I saw this somewhere and decided to use it
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   60 * time.Second,
			KeepAlive: 60 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 10 * time.Second,
	}
	proxy := &httputil.ReverseProxy{
		Director:  director,
		Transport: transport,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	})

	log.Printf("Proxy server running on port 9001!!\n")

	server := http.ListenAndServe(":9001", nil)
	log.Fatal(server, "\n")
}
