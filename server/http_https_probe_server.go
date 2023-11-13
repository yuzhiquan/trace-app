package server

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"time"
)

func DoHTTPSRequest() {
	log.Println("https request is begin.")
	go func() {
		index := 0
		for {
			tr := &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}
			client := &http.Client{Transport: tr}
			_, err := client.Get("https://httpbin.org/")
			if err != nil {
				fmt.Println(err)
			}
			//_, err = ioutil.ReadAll(res.Body)
			//if err != nil {
			//	fmt.Printf("client: could not read response body: %s\n", err)
			//}
			//fmt.Printf("client: response body: %s\n", resBody)
			fmt.Printf("request 0%d @ %s\n", index, time.Now())
			//time.Sleep(5 * time.)
			index += 1
		}
	}()
}

func DoHTTPRequest(url string) {
	log.Printf("http request => %s begin.", url)
	go func() {
		index := 0
		for {
			tr := &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}
			client := &http.Client{Transport: tr}
			_, err := client.Get("https://httpbin.org/")
			if err != nil {
				fmt.Println(err)
			}
			//_, err = ioutil.ReadAll(res.Body)
			//if err != nil {
			//	fmt.Printf("client: could not read response body: %s\n", err)
			//}
			//fmt.Printf("client: response body: %s\n", resBody)
			fmt.Printf("request 0%d @ %s\n", index, time.Now())
			//time.Sleep(5 * time.)
			index += 1
		}
	}()
}
