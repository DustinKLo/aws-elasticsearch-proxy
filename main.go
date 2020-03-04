package main

import (
	"net/http"

	"github.com/hysds/aws-elasticsearch-proxy/configs"
	L "github.com/hysds/aws-elasticsearch-proxy/logger"
	proxy "github.com/hysds/aws-elasticsearch-proxy/reverseproxy"
)

func main() {
	host := configs.Host
	if host == "" {
		L.Logging.Warning("host not found in settings.yml, defaulting to http://localhost:9200")
		host = "http://localhost:9200"
	}

	reverseProxy := proxy.AwsEsReverseProxy(host)

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
