package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/lib/pq" // pg connector
)

//PostgresConfig is db config
type PostgresConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
	Host     string `json:"host"`
	Port     uint64 `json:"port"`
}

// FetchConfig is fetcher config
type FetchConfig struct {
	DbConfig string `json:"raw_storage_config"`
	BaseURL  string `json:"base_url"`
}

// Fetcher downloads raw data and stores it in db
type Fetcher struct {
	Config     FetchConfig
	stashes    chan (PublicStashesResponse)
	nextID     chan (string)
	done       chan (int64)
	lastNextID string
}

func (f *Fetcher) fetch() {
	fmt.Println("start fetch")
	for {
		resp, err := http.Get(f.Config.BaseURL + f.lastNextID)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		var msg PublicStashesResponse
		err = json.NewDecoder(resp.Body).Decode(&msg)
		if err != nil {
			log.Fatal(err)
		}
		f.stashes <- msg
		f.lastNextID = <-f.nextID
	}
}

func (f *Fetcher) storeData() {
	raw, err := ioutil.ReadFile(f.Config.DbConfig)
	if err != nil {
		log.Panic(err)
	}
	var dbConfig PostgresConfig
	err = json.Unmarshal(raw, &dbConfig)
	if err != nil {
		log.Panic(err)
	}

	r := strings.NewReplacer(
		"$(user)$", dbConfig.User,
		"$(password)$", dbConfig.Password,
		"$(host)$", dbConfig.Host,
		"$(port)$", strconv.Itoa(int(dbConfig.Port)),
		"$(db)$", dbConfig.Database,
	)

	connStr := "postgres://$(user)$:$(password)$@$(host)$:$(port)$/$(db)$?sslmode=disable"
	res := r.Replace(connStr)
	fmt.Println(res)
	db, err := sql.Open("postgres", res)
	if err != nil {
		log.Fatal(err)
	}
	query := "insert into tbl_raw_stash_data(stash_id, raw_data) values ($1, $2)"
	for {
		msg := <-f.stashes
		fmt.Println(msg.NextChangeID)

		msgStr, err := json.Marshal(msg)
		if err != nil {
			log.Fatal(err)

		}
		_, err = db.Exec(query, msg.NextChangeID, string(msgStr))
		if err != nil {
			log.Fatal(err)

		}
		f.nextID <- msg.NextChangeID
	}

}

//FetchData stores data in db
func FetchData(configPath string) {
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
	fetcher := Fetcher{
		config,
		make(chan PublicStashesResponse),
		make(chan string),
		make(chan int64),
		"202645757-210123687-198063514-227613078-213939523", // TODO store last next id in database
	}
	go fetcher.fetch()
	go fetcher.storeData()

	<-fetcher.done
}
