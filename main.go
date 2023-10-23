package main

import (
	"context"
	"flag"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"time"
)

func main() {
	nodeIp := os.Getenv("NODE_IP")
	nodeZone := os.Getenv("NODE_ZONE")
	nodeRegion := os.Getenv("NODE_REGION")
	agentUrl := flag.String("zipkin", fmt.Sprintf("http://%s:9411", nodeIp), "zipkin url")
	flag.Parse()
	log.Printf("NODE_IP: %s, agentUtrl:%+v", nodeIp, agentUrl)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	shutdown, err := initTracer(*agentUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := shutdown(ctx); err != nil {
			log.Fatal("failed to shutdown TracerProvider: %w", err)
		}
	}()

	tr := otel.GetTracerProvider().Tracer("component-main")

	// define origin server URL
	originServerURL, err := url.Parse("http://127.0.0.1:12345")
	if err != nil {
		log.Fatal("invalid origin server URL")
	}

	reverseProxy := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		_, span := tr.Start(ctx, "trace-app", trace.WithSpanKind(trace.SpanKindServer))
		defer span.End()
		fmt.Printf("[reverse proxy server] received request at: %s\n", time.Now())

		// set req Host, URL and Request URI to forward a request to the origin server
		req.Host = originServerURL.Host
		req.URL.Host = originServerURL.Host
		req.URL.Scheme = originServerURL.Scheme
		req.RequestURI = ""
		req.Host = "www.baidu.com:80"

		//req.Header.Add("host", "www.baidu.com:80")

		log.Printf("req is :%+v", req)
		// save the response from the origin server
		originServerResponse, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Printf("got err in do request to originserver:%+v", err)
			rw.WriteHeader(http.StatusInternalServerError)
			_, _ = fmt.Fprint(rw, err)
			return
		}

		// return response to the client
		rw.WriteHeader(http.StatusOK)

		io.Copy(rw, originServerResponse.Body)

		log.Printf("originServerResponse:%+v", originServerResponse.Body)
		for k, v := range originServerResponse.Trailer {
			for _, value := range v {
				rw.Header().Set(k, value)
			}
		}

		rw.Header().Set("content-type", "text/html; charset=utf-8")

	})

	go http.ListenAndServe(":8080", reverseProxy)

	http.HandleFunc("/region", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, fmt.Sprintf("request to region: %s", nodeRegion))
	})

	http.HandleFunc("/zone", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, fmt.Sprintf("request to zone: %s", nodeZone))
	})

	http.HandleFunc("/rezone", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, fmt.Sprintf("request to region: %s zone:%s", nodeRegion, nodeRegion))
	})

	http.ListenAndServe(":8090", nil)
}
