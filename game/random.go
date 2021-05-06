package game

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const RANDOM_URL = "https://codechallenge.boohma.com/random"

var client *http.Client

func init() {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 100
	t.MaxConnsPerHost = 100
	t.MaxIdleConnsPerHost = 100
	client = &http.Client{
		Timeout:   time.Second * 10,
		Transport: t,
	}
}

type randomResponse struct {
	Value int `json:"random_number"`
}

func GetRandomNumber() int {
	// client := &http.Client{
	// 	Timeout: time.Second * 10,
	// }
	req, err := http.NewRequest("GET", RANDOM_URL, nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	var randomResponse randomResponse
	json.Unmarshal(bodyBytes, &randomResponse)

	return randomResponse.Value
}
