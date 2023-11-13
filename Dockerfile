#build stage
#FROM --platform=linux/amd64 golang AS builder
##RUN apk add --no-cache git
#WORKDIR /go/src/app
#COPY . .
#RUN go get -d -v ./...
## RUN mkdir -p /go/bin/
#RUN env GOOS=linux GOARCH=amd64 go build  -v ./...

##final stage
FROM --platform=linux/amd64 debian:buster
#RUN apt-get install ca-certificates
#FROM --platform=linux/amd64 alpine:latest
# RUN apk --no-cache add ca-certificates

COPY ./trace-app /app/demo
#RUN apk --no-cache add ca-certificates
#COPY --from=builder /go/src/app/trace-app /app/demo
#ENTRYPOINT ["/app/demo"]