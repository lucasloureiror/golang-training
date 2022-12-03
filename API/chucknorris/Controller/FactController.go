package Controller

import (
	"chucknorris/model"
	"encoding/json"
	"fmt"
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

func RandomFactWithCategory(category string) model.Fact {
	resp, err := http.Get("https://api.chucknorris.io/jokes/random?category=" + category)

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var output model.Fact

	json.Unmarshal([]byte(responseData), &output)

	fmt.Println(output)

	return output

}
