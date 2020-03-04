FROM golang:1.14 as builder

WORKDIR /app

COPY . .

RUN go build .

FROM golang:1.14

LABEL maintainer="dustin.k.lo@nasa.jpl.gov"

COPY --from=builder /app/aws-elasticsearch-proxy /aws-elasticsearch-proxy

EXPOSE 9001

CMD ["/aws-elasticsearch-proxy"]
