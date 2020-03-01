package utils

import (
	"os"

	L "github.com/hysds/aws-elasticsearch-proxy/logger"
)

func GetAwsRegion() string {
	region := os.Getenv("AWS_REGION") // "us-west-1"
	if region == "" {
		errMsg := "AWS_REGION environment variable not provided"
		L.Logging.Fatal(errMsg)
		panic(errMsg)
	}
	return region
}

func GetEsEndpoint() string {
	host := os.Getenv("HOST")
	if host == "" {
		L.Logging.Warning("HOST environment variable not found, defaulting to localhost:9200")
		return "http://localhost:9200"
	}
	return host
}
