# syntax=docker/dockerfile:1

FROM golang:1.25.5 AS builder

# Set destination for COPY
WORKDIR /app
COPY go.work go.work.sum ./
COPY backend/ ./backend/
COPY worker/ ./worker/
COPY pkg/ ./pkg/
RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct && \
    go mod download
# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o worker ./worker

FROM alpine:3
RUN apk add --no-cache pngquant jpegoptim
COPY --from=builder /app/worker /bin/worker
ENTRYPOINT ["/bin/worker"]