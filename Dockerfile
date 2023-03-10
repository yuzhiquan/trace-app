#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
# RUN mkdir -p /go/bin/
RUN go build  -v ./...

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/src/app/trace-app /app/demo
ENTRYPOINT ["/app/demo"]