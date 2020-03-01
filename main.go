package main

import (
	"net/http"

	creds "github.com/hysds/aws-elasticsearch-proxy/awscredentials"
	L "github.com/hysds/aws-elasticsearch-proxy/logger"
	proxy "github.com/hysds/aws-elasticsearch-proxy/reverseproxy"
	utils "github.com/hysds/aws-elasticsearch-proxy/utils"
)

func main() {
	service := "es"
	region := utils.GetAwsRegion()
	signer := creds.GetAwsSigner()

	host := utils.GetEsEndpoint() // use environment variables instead
	reverseProxy := proxy.AwsEsReverseProxy(host, signer, region, service)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.Header().Add("Connection", "keep-alive") //handling preflight
			w.Header().Add("Access-Control-Allow-Origin", "*")
			w.Header().Add("Access-Control-Allow-Methods", "POST, OPTIONS, GET, DELETE, PUT")
			w.Header().Add("Access-Control-Allow-Headers", "content-type")
		} else {
			reverseProxy.ServeHTTP(w, r)
		}
	})

	port := ":9001" // use environment variable
	L.Logging.Info("Proxy server running on port", port)
	server := http.ListenAndServe(port, nil)
	L.Logging.Fatal(server)
}
