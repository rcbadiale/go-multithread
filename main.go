package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

/* Make a GET request and write the response as a string to the channel. */
func getURL(url string, ch chan<- string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	d, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	ch <- string(d)
}

func main() {
	args := os.Args
	if len(args) != 2 {
		log.Fatal("Invalid command.\n  Usage: go run main.go <cep>")
	}
	cep := args[1]

	ch1 := make(chan string)
	url1 := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	go getURL(url1, ch1)

	ch2 := make(chan string)
	url2 := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)
	go getURL(url2, ch2)

	select {
	case data := <-ch1:
		fmt.Printf("URL: %s\nResponse: %s", url1, data)
	case data := <-ch2:
		fmt.Printf("URL: %s\nResponse: %s", url2, data)
	case <-time.After(1 * time.Second):
		fmt.Println("timed out after 1s")
	}
}
