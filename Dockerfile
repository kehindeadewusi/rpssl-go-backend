FROM golang:1.16-alpine as builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o rpssl .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist
RUN cp /build/rpssl .

#Good man! https://doc.xuwenliang.com/docs/go/2056
RUN echo "nobody:x:65534:65534:Nobody:/:" > /etc_passwd

# Production BASE image
FROM scratch

COPY --from=builder /dist/rpssl /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc_passwd /etc/passwd

# Run as the new non-root by default
USER nobody

ENTRYPOINT ["/rpssl"]