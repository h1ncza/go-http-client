package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	URL string
)

func init() {
	flag.StringVar(&URL, "h", "", "URL to connect to")
}

func main() {
	flag.Parse()

	if len(URL) == 0 {
		fmt.Println("Usage: http-goclient -h <HOST>")
		// flag.PrintDefaults()
		os.Exit(1)
	} else if strings.HasPrefix(URL, "http://") || strings.HasPrefix(URL, "https://") {
		URL = URL
	} else {
		URL = "http://" + URL
	}

	resp, err := http.Get(URL)
	if err != nil {
		fmt.Println("Connection refused")
		os.Exit(1)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)
	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(string(respByte))
}
