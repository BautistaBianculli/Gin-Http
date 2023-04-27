package api

import (
	"GORUTINE/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

var client *http.Client

func GetMoovies(mChan chan []models.Moovies, i int, wg *sync.WaitGroup) {
	url, ok := os.LookupEnv("API_URL")
	if !ok {
		panic("API_URL NOT SET")
	}
	key, ok := os.LookupEnv("API_KEY")
	if !ok {
		panic("APIC_KEY NOT SET")
	}
	url = fmt.Sprintf("%s%d%s", url, i, key)
	client = &http.Client{Timeout: 10 * time.Second}
	response, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	var data models.Moovies
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		panic(err)
	}
	response.Body.Close()
	movies := []models.Moovies{data}
	mChan <- movies
	wg.Done()

}
