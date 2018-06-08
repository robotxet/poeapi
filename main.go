package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/robotxet/poeapi/api"
)

func main() {
	resp, err := http.Get("http://api.pathofexile.com/public-stash-tabs/?id=202645757-210123687-198063514-227613078-213939523")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()
	var msg api.PublicStashes
	err = json.NewDecoder(resp.Body).Decode(&msg)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
