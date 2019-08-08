FROM golang:1.12.1-alpine3.9 AS based_builder

# Install git, curl and openssh
RUN apk add --no-cache git curl openssh

# Download go modules as pre-caches for derivative build
WORKDIR /src/go
COPY go.mod go.sum ./
RUN go mod download

# =============================================================================
FROM based_builder AS builder

# Receive argument from command when building a service
ARG service

# Copy the code from the host and compile it
WORKDIR /src/go
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app .

# =============================================================================
FROM alpine:3.9.2
RUN apk update && \
   apk add curl ca-certificates && \
   rm -rf /var/cache/apk/** && \
   update-ca-certificates

COPY --from=builder /app ./
ENTRYPOINT ["./app"]
