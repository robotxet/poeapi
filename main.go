package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Socket struct {
	Group  int64       `json:"group"`
	Attr   interface{} `json:"attr"`
	Colour string      `json:"sColour"`
}

type AdditionaItemProperty struct {
	Name        string        `json:"name"`
	Values      []interface{} `json:"values"`
	DisplayMode int64         `json:"displayMode"`
	Progress    float64       `json:"progress"`
	Type        int64         `json:"type"`
}

type Property struct {
	Name        string          `json:"name"`
	Values      [][]interface{} `json:"values"`
	DisplayMode int64           `json:"displayMode"`
}
type ItemCategory struct {
	Info map[string][]string
}

type Item struct {
	Verified              bool                    `json:"verified"`
	W                     int64                   `json:"w"`
	H                     int64                   `json:"h"`
	ID                    string                  `json:"id"`
	Ilvl                  int64                   `json:"ilvl"`
	Icon                  string                  `json:"icon"`
	League                string                  `json:"league"`
	Sockets               []Socket                `json:"sockets"`
	Name                  string                  `json:"name"`
	TypeLine              string                  `json:"typeLine"`
	Identified            bool                    `json:"identified"`
	Note                  string                  `json:"note"`
	Properties            []Property              `json:"properties"`
	AdditionalPropreties  []AdditionaItemProperty `json:"additionalProperties"`
	Requirements          []Property              `json:"requirements"`
	NextLevelRequirements []Property              `json:"nextLevelRequirements"`
	SecDescrText          string                  `json:"secDescrText"`
	ExplicitMods          []string                `json:"explicitMods"`
	FlavourText           []string                `json:"flavourText"`
	FrameType             int64                   `json:"frameType"`
	Category              ItemCategory            `json:"category"`
	X                     int64                   `json:"x"`
	Y                     int64                   `json:"y"`
	InventoryID           string                  `json:"inventoryId"`
}

type Stash struct {
	AccountName       string `json:"accountName"`
	LastCharacterName string `json:"lastCharacterName"`
	ID                string `json:"id"`
	StashName         string `json:"stash"`
	StashType         string `json:"stashType"`
	Items             []Item `json:"items"`
	Public            bool   `json:"public"`
}

type PoeResponse struct {
	NextChangeID string  `json:"next_change_id"`
	Stashes      []Stash `json:"stashes"`
}

func main() {
	fmt.Println("test")
	resp, err := http.Get("http://api.pathofexile.com/public-stash-tabs/?id=202645757-210123687-198063514-227613078-213939523")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var msg PoeResponse
	err = json.Unmarshal(body, &msg)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
