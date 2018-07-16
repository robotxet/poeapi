package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/olivere/elastic"
)

//https://github.com/olivere/elastic/wiki/QueryDSL

//CheckElastic just test for elastic connector
func CheckElastic() {
	ctx := context.Background()
	client, err := elastic.NewClient()
	if err != nil {
		log.Panic(err)
	}

	info, code, err := client.Ping("http://127.0.0.1:9200").Do(ctx)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
}

//ElasticConfig is config for elasticsearch
type ElasticConfig struct {
	Host string `json:"host"`
	Port uint64 `json:"port"`
}

//ElasticProcessor for elasticsearch
type ElasticProcessor struct {
	Config     FetchConfig
	stashes    chan (PublicStashesResponse)
	nextID     chan (string)
	done       chan (int64)
	lastNextID string
}

func (es *ElasticProcessor) pullStashes() {
	fmt.Println("start fetch")
	for {
		resp, err := http.Get(es.Config.BaseURL + es.lastNextID)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		var msg PublicStashesResponse
		err = json.NewDecoder(resp.Body).Decode(&msg)
		if err != nil {
			log.Fatal(err)
		}
		es.stashes <- msg
		es.lastNextID = <-es.nextID
		log.Println(es.lastNextID)
	}
}

func (es *ElasticProcessor) storeData() {
	ctx := context.Background()
	client, err := elastic.NewClient()
	if err != nil {
		log.Panic(err)
	}

	_, _, err = client.Ping("http://127.0.0.1:9200").Do(ctx)
	if err != nil {
		log.Panic(err)
	}

	// exists, err := client.IndexExists("public_stashes").Do(ctx)
	// if err != nil {
	// 	log.Panic(err)
	// }
	// if !exists {
	// 	var stash PublicStashesResponse
	// 	createIndex, err := client.CreateIndex("public_stashes").BodyJson(stash).Do(ctx)
	// 	if err != nil {
	// 		log.Panic(err)
	// 	}
	// 	if !createIndex.Acknowledged {
	// 		log.Println("now acknowledged index")
	// 	}
	// }
	for {
		msg := <-es.stashes
		_, err = client.Index().Index("public_stashes").Type("public_stash").BodyJson(msg).Do(ctx)
		if err != nil {
			log.Println("can't index")
			log.Panic(err)
		}
		fmt.Println(msg.NextChangeID)
		es.nextID <- msg.NextChangeID
		es.done <- 1
	}
}

func (es *ElasticProcessor) searchData() {
	ctx := context.Background()
	client, err := elastic.NewClient()
	if err != nil {
		log.Panic(err)
	}

	_, _, err = client.Ping("http://127.0.0.1:9200").Do(ctx)
	if err != nil {
		log.Panic(err)
	}
	searchQuery := elastic.NewNestedQuery("stashes", elastic.NewMatchQuery("stashes.stashType", "MapStash"))
	query := elastic.NewBoolQuery()
	query.Filter(searchQuery)
	searchResult, err := client.Search("public_stashes").
		Type("public_stash").
		Query(elastic.NewMatchAllQuery()).
		Do(ctx)
	src, err := query.Source()
	fmt.Println(src)
	if err != nil {
		// Handle error
		log.Panic(err)
	}
	for _, item := range searchResult.Hits.Hits {
		fmt.Println("test")
		var resStash PublicStashesResponse
		if err := json.Unmarshal(*item.Source, &resStash); err != nil {
			log.Panic(err)
		}
		fmt.Println(resStash.NextChangeID)
	}

	es.done <- 1

}

//ElasticTest is test
func ElasticTest(configPath string) {
	configStr, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Println(err.Error())
		return
		//log.Panic(err.Error())
	}
	var config FetchConfig
	err = json.Unmarshal(configStr, &config)
	if err != nil {
		fmt.Println(err.Error())
		return
		//log.Panic(err.Error())
	}
	fetcher := ElasticProcessor{
		config,
		make(chan PublicStashesResponse, 2),
		make(chan string, 2),
		make(chan int64),
		"202645757-210123687-198063514-227613078-213939523", // TODO store last next id in database
	}
	// go fetcher.pullStashes()
	// go fetcher.storeData()
	go fetcher.searchData()

	<-fetcher.done

}
