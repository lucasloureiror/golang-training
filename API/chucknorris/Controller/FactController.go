package Controller

import (
	"chucknorris/model"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func RandomFact() model.Fact {

	resp, err := http.Get("https://api.chucknorris.io/jokes/random")

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var output model.Fact

	json.Unmarshal([]byte(responseData), &output)

	return output

}
