package api

// FetchConfig is fetcher config
type FetchConfig struct {
	DbConfig string `json:"raw_storage_path"`
	BaseURL  string `json:"base_url"`
}

// Fetcher downloads raw data and stores it in db
type Fetcher struct {
	Config FetchConfig
	Chan   chan (PublicStashesResponse)
}

func (*Fetcher) fetch(nextID string) {

}

func (*Fetcher) storeData() {

}
