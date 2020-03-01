## Set AWS Credentials as environment variable

```
# make sure there is a $HOME/.aws/credentials file
export AWS_REGION=us-west-1

# other environment variables
export LOG_FILE_PATH=/path/to/log/file  # will default to stdout if not found
export LOG_LEVEL=debug  # (warn, warning, debug, error)
export HOST=http://localhost:9200  # Elasticsearch endpoint
```

## Run app
```
go run main.go
```

## Example of signed request
```
2020/02/26 21:33:52 Content-Length [357]
2020/02/26 21:33:52 Accept [application/json]
2020/02/26 21:33:52 Sec-Fetch-Mode [cors]
2020/02/26 21:33:52 Accept-Encoding [gzip, deflate, br]
2020/02/26 21:33:52 Accept-Language [en-US,en;q=0.9]
2020/02/26 21:33:52 User-Agent [Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36]
2020/02/26 21:33:52 Sec-Fetch-Site [same-site]
2020/02/26 21:33:52 Content-Type [application/x-ndjson]
2020/02/26 21:33:52 X-Forwarded-Host [localhost:9001]
2020/02/26 21:33:52 X-Destination-Host [localhost:9200]
2020/02/26 21:33:52 Origin [http://localhost:8080]
2020/02/26 21:33:52 Referer [http://localhost:8080/tosca]
2020/02/26 21:33:52 Connection [keep-alive]
2020/02/26 21:33:52 X-Amz-Date [20200227T053352Z]
2020/02/26 21:33:52 X-Amz-Security-Token [################################]
2020/02/26 21:33:52 Authorization [AWS4-HMAC-SHA256 Credential=#####################/20200227/us-west-1/es/aws4_request, SignedHeaders=accept;accept-encoding;accept-language;connection;content-length;content-type;host;origin;referer;sec-fetch-mode;sec-fetch-site;x-amz-date;x-amz-security-token;x-destination-host;x-forwarded-host, Signature=########################]
```