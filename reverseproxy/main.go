package reverseproxy

import (
	"bytes"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
	L "github.com/hysds/aws-elasticsearch-proxy/logger"
)

func AwsEsReverseProxy(host string, signer *v4.Signer, region string, service string) *httputil.ReverseProxy {
	origin, _ := url.Parse(host)
	reverseProxy := httputil.NewSingleHostReverseProxy(origin)

	reverseProxy.Director = func(req *http.Request) {
		req.Header.Set("Accept", "*/*")
		req.Header.Add("X-Forwarded-Host", req.Host)
		req.Header.Add("X-Origin-Host", origin.Host)
		req.URL.Scheme = "http"
		req.URL.Host = origin.Host

		if req.Body == nil { // for GET, DELETE requests
			signer.Sign(req, nil, service, region, time.Now()) // signing request
			L.Logging.Info(req)
			return
		}

		body, err := ioutil.ReadAll(req.Body) // reading request body for AWS signature
		if err != nil {
			L.Logging.Error("Error reading body:", err)
			return
		}

		// Restore io.ReadCloser to its original state (request body can only be read once)
		req.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		newBody := bytes.NewReader(body)                       // converting to type io.Reader for AWS signature
		signer.Sign(req, newBody, service, region, time.Now()) // signing request
		L.Logging.Info(req)
	}

	reverseProxy.ErrorHandler = func(w http.ResponseWriter, req *http.Request, err error) {
		L.Logging.Error(err)
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}

	// not sure I saw this somewhere and decided to use it
	reverseProxy.Transport = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   60 * time.Second,
			KeepAlive: 60 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	return reverseProxy
}
