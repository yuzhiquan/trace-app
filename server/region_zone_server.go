package server

import (
	"fmt"
	"net/http"
	"os"
)

func RunRegionZoneServer() {
	nodeIp := os.Getenv("NODE_IP")
	nodeZone := os.Getenv("NODE_ZONE")
	nodeRegion := os.Getenv("NODE_REGION")

	http.HandleFunc("/region", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, fmt.Sprintf("request to region: %s nodeIP: %s\n", nodeRegion, nodeIp))
	})

	http.HandleFunc("/zone", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, fmt.Sprintf("request to zone: %s nodeIP: %s\n", nodeZone, nodeIp))
	})

	http.HandleFunc("/rezone", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, fmt.Sprintf("request to region: %s zone:%s nodeIP: %s\n", nodeRegion, nodeZone, nodeIp))
	})

	http.ListenAndServe(":8090", nil)
}
