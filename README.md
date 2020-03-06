## Loading configurations from settings.yml
```
# settings.yml
# ERROR: 40
# WARNING: 30
# INFO: 20
# DEBUG: 10
log_level: 10
log_file_path: /path/to/log/file.log

host: http://localhost:9200
http_scheme: http
verify_ssl: false

service: es
aws_region: us-west-2
```

## Run application
```
go run main.go
```

## Dockerize application
```
# docker build
docker build -t aws-elasticsearch-proxy:latest .

# docker run with environment variable file
docker run \
    -v ~/.aws:/root/.aws \
    -p 9001:9001 \
    aws-elasticsearch-proxy:latest
```

## Example of signed request
```
[2020-03-05T20:07:23-08:00]  INFO POST /_all/_search HTTP/1.1
Host: localhost:9200
User-Agent: Go-http-client/1.1
Content-Length: 28
Authorization: AWS4-HMAC-SHA256 Credential=#####################/20200306/us-west-2/es/aws4_request, SignedHeaders=content-type;host;x-amz-date;x-amz-security-token, Signature=###############################################################
Content-Type: application/json
X-Amz-Date: 20200306T040723Z
X-Amz-Security-Token: #####################//////////#####################==
Accept-Encoding: gzip

{"query": {"match_all": {}}}
```