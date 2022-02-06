package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	host string
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: gocurl <host>")
		os.Exit(1)
	} else {
		host = os.Args[1]
	}

	if strings.HasPrefix(host, "http://") || strings.HasPrefix(host, "https://") {
		host = host
	} else {
		host = "http://" + host
	}

	resp, err := http.Get(host)
	if err != nil {
		fmt.Println("Connection refused")
		os.Exit(1)
	}
	defer resp.Body.Close()

	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(string(respByte))
}
