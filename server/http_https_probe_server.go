package server

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"time"
)

func DoHTTPSRequest(url string, interval int) {
	log.Println("https request is begin.")
	if url == "" {
		url = "https://httpbin.org/"
	}
	go func() {
		for {
			time.Sleep(time.Duration(interval) * time.Second)

			tr := &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}
			client := &http.Client{Transport: tr}
			resp, err := client.Get(url)
			if err != nil {
				fmt.Println(err)
				continue
			}
			//_, err = ioutil.ReadAll(res.Body)
			//if err != nil {
			//	fmt.Printf("client: could not read response body: %s\n", err)
			//}
			//fmt.Printf("client: response body: %s\n", resBody)
			fmt.Printf("https request => %s response code: %d  @ %s\n", url, resp.StatusCode, time.Now())
		}
	}()
}

func DoHTTPRequest(url string, interval int) {
	if url == "" {
		url = "http://httpbin.org"
	}
	log.Printf("http request => %s begin.", url)
	go func() {
		for {
			time.Sleep(time.Duration(interval) * time.Second)

			res, err := http.Get(url)
			if err != nil {
				fmt.Printf("error making http request: %s\n", err)
				continue
			}
			fmt.Printf("http request=> %s response code: %d @%s\n", url, res.StatusCode, time.Now())
		}
	}()
}
