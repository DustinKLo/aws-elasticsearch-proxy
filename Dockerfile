# Build Go binary in builder
FROM golang:1.14 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build .

# or other base image you want
FROM golang:1.14

# Add Maintainer Info
LABEL maintainer="dustin.k.lo@nasa.jpl.gov"

COPY --from=builder /app/aws-elasticsearch-proxy /aws-elasticsearch-proxy

# Expose port 8080 to the outside world
EXPOSE 9001

# Command to run the executable
CMD ["/aws-elasticsearch-proxy"]
