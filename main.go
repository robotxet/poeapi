package main

import (
	"flag"
	"fmt"

	"poeapi/api"
)

func main() {
	config := flag.String("config", "", "fetcher config")
	flag.Parse()
	fmt.Println("running")
	api.FetchData(*config)
	fmt.Println("end")
}
