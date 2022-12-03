package Controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Fact struct {
	Value string `json:"value"`
	Id    string `json:"id"`
}

func RandomFact() Fact {

	resp, err := http.Get("https://api.chucknorris.io/jokes/random")

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var output Fact

	json.Unmarshal([]byte(responseData), &output)

	return output

}
