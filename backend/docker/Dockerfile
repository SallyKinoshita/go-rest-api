FROM golang:1.23.3 AS builder

ARG TARGETOS=linux
ARG TARGETARCH=amd64

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -a -o backend ./cmd/server/main.go

FROM alpine:latest

ENV TZ=Asia/Tokyo

WORKDIR /
COPY --from=builder /app/backend /backend
USER 65532:65532

ENTRYPOINT ["/backend"]
