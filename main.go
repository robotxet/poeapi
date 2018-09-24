package main

import (
	"flag"
	"fmt"
	"github.com/robotxet/poeapi/api"
)

func main() {
	config := flag.String("config", "", "fetcher config")
	flag.Parse()
	fmt.Println("running")
	api.ElasticTest(*config)
	fmt.Println("end")

}
