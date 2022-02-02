package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

// UrlParsing shows some examples of working with URLs using
// net/url package
func UrlParsing() {
	fmt.Println("========== Exploring net/url ==========")
	const httpbinPost = "https://httpbin.org/post?method=TestEcho&msg=Hello&count=1"
	u, err := url.Parse(httpbinPost)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(">>> Checking out URL fields and methods")
	fmt.Printf("u: %v\n", u)
	fmt.Printf("u.Host: %v\n", u.Host)
	fmt.Printf("u.Path: %v\n", u.Path)
	fmt.Printf("u.Scheme: %v\n", u.Scheme)
	fmt.Printf("u.Hostname(): %v\n", u.Hostname())
	fmt.Printf("u.Port(): %v\n", u.Port())

	// Query() returns the query data in key-value pairs
	values := u.Query()
	for k, v := range values {
		fmt.Println(k, ":", v)
	}
	fmt.Printf("u.String(): %v\n", u.String())
}

// PostSomething shows some example calls of using http.Post
// It sends POST requests to httpbin with some dummy args
// and data in json format
func PostSomething() {
	fmt.Println("========== Exploring http.Post() functionalities ==========")

	url := "https://httpbin.org/post?method=TestEcho&msg=Hello&count=1"
	resp, err := http.Post(url, "", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Printf("url: %v\n", url)
	fmt.Printf("resp.Status: %v\n", resp.Status)
	fmt.Printf("resp.StatusCode: %v\n", resp.StatusCode)
	replyData, _ := io.ReadAll(resp.Body)
	fmt.Printf("replyData: %s\n", replyData)

	// Encoding a struct into json and passing that in the data
	// field of the post request

	type Payload struct {
		Method  string `json:"method"`
		Message string `json:"messasge"`
		Count   int    `json:"count"`
	}

	p1 := Payload{"Echo", "Muhahaha", 2}
	payload, err := json.Marshal(p1)
	if err != nil {
		log.Fatal(err)
	}

	url = "https://httpbin.org/post"
	resp, err = http.Post(url, "application/json", bytes.NewReader(payload))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Printf("url: %v\n", url)
	fmt.Printf("resp.Status: %v\n", resp.Status)
	fmt.Printf("resp.StatusCode: %v\n", resp.StatusCode)
	replyData, _ = io.ReadAll(resp.Body)
	fmt.Printf("replyData: %s\n", replyData)
}
