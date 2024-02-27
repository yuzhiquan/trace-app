package server

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func RunRegionZoneServer(random bool) {
	nodeIp := os.Getenv("NODE_IP")
	nodeZone := os.Getenv("NODE_ZONE")
	nodeRegion := os.Getenv("NODE_REGION")
	randomFail := false
	randomFailInt := 0

	http.HandleFunc("/region", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, fmt.Sprintf("request to region: %s nodeIP: %s\n", nodeRegion, nodeIp))
	})

	http.HandleFunc("/zone", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, fmt.Sprintf("request to zone: %s nodeIP: %s\n", nodeZone, nodeIp))
	})

	http.HandleFunc("/rezone", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, fmt.Sprintf("request to region: %s zone:%s nodeIP: %s\n", nodeRegion, nodeZone, nodeIp))
	})
	http.HandleFunc("/randomfail/", func(writer http.ResponseWriter, request *http.Request) {
		randomFail = true
		id := strings.TrimPrefix(request.URL.Path, "/randomfail/")
		i, err := strconv.Atoi(id)
		if err != nil {
			randomFailInt = 2
		}
		randomFailInt = i
		fmt.Fprintf(writer, fmt.Sprintf("set randomfail: %t, randomfail int:%d \n", randomFail, randomFailInt))
	})

	http.HandleFunc("/status/", func(writer http.ResponseWriter, request *http.Request) {
		id := strings.TrimPrefix(request.URL.Path, "/status/")
		i, err := strconv.Atoi(id)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, fmt.Sprintf("can not parse the status code: %s\n", id))
			return
		}
		if rand.Intn(100)%3 == 0 && random {
			writer.WriteHeader(http.StatusOK)
			fmt.Fprintf(writer, fmt.Sprintf("parse the status code: %s but return 200\n", id))
			return
		}

		if randomFail && rand.Intn(100)%randomFailInt == 0 {
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, fmt.Sprintf("parse the status code: %s but return 500\n", id))
			return
		}
		writer.WriteHeader(i)
		fmt.Fprintf(writer, fmt.Sprintf("parse the status code: %s\n", id))
	})

	http.ListenAndServe(":8090", nil)
}
