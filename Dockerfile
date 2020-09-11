FROM golang:alpine AS builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 

# in case you build image behind firewall
# ENV GOPROXY=https://goproxy.cn

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o main .
RUN go build -o ./gateway/main ./gateway/main.go

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/main .
RUN mkdir gateway
RUN cp /build/gateway/main ./gateway

# Build a small image
FROM scratch AS grpc

EXPOSE 9192

COPY --from=builder /dist/main /

# Command to run
ENTRYPOINT ["/main"]

FROM scratch AS gateway

EXPOSE 8080

COPY --from=builder /dist/gateway/main /

ENTRYPOINT [ "/main" ]