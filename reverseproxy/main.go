package reverseproxy

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	creds "github.com/hysds/aws-elasticsearch-proxy/awscredentials"
	"github.com/hysds/aws-elasticsearch-proxy/configs"
	L "github.com/hysds/aws-elasticsearch-proxy/logger"
)

func AwsEsReverseProxy(host string) *httputil.ReverseProxy {
	origin, _ := url.Parse(host)
	reverseProxy := httputil.NewSingleHostReverseProxy(origin)

	service := configs.Service
	region := configs.AWSRegion
	signer := creds.GetSigner()

	reverseProxy.Director = func(req *http.Request) {
		req.Header.Del("Accept")
		req.Header.Del("Content-Length")
		req.Header.Del("X-Forwarded-For")
		req.Header.Del("User-Agent")
		req.Header.Del("Connection")

		req.Host = origin.Host
		req.URL.Host = origin.Host
		req.URL.Scheme = configs.HttpScheme

		if req.Body == nil { // for GET, DELETE requests
			signer.Sign(req, nil, service, region, time.Now()) // signing request
		} else {
			body, err := ioutil.ReadAll(req.Body) // reading request body for AWS signature
			if err != nil {
				L.Logging.Error("Error reading body:", err)
				return
			}
			req.Body = ioutil.NopCloser(bytes.NewBuffer(body))     // restore io.ReadCloser to its original state (can only be read once)
			newBody := bytes.NewReader(body)                       // converting to type io.Reader for AWS signature
			signer.Sign(req, newBody, service, region, time.Now()) // signing request
		}

		if configs.LogLevel <= 20 {
			dump, err := httputil.DumpRequestOut(req, true)
			if err != nil {
				L.Logging.Debug(err)
				return
			}
			L.Logging.Info(string(dump))
		}
	}

	reverseProxy.ErrorHandler = func(w http.ResponseWriter, req *http.Request, err error) {
		L.Logging.Error(err)
		w.WriteHeader(503)
		w.Write([]byte(err.Error()))
	}

	proxyTransport := &http.Transport{ // not sure I saw this somewhere and decided to use it
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
	if configs.VerifySSL == false {
		proxyTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true} // verify SSL
	}
	reverseProxy.Transport = proxyTransport

	return reverseProxy
}
