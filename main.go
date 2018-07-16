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
	api.ElasticTest(*config)
	fmt.Println("end")

}
