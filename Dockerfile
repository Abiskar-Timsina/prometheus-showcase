FROM golang:bookworm AS builder
WORKDIR app/
COPY ./go.mod .
RUN go mod download
COPY . ./
RUN go build -o main ./cmd/app/main.go

FROM alpine:latest
WORKDIR application/
RUN apk add libc6-compat # Required for alpine to work with go (this depends on the host machine mostly)
COPY --from=builder /go/app/main /application/
ENTRYPOINT ["/application/main"]
