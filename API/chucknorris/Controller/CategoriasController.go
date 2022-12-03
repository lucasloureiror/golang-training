package Controller

import (
	"io/ioutil"
	"log"
	"net/http"
)

func Categorias() string {

	resp, err := http.Get("https://api.chucknorris.io/jokes/categories")

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	output := string(responseData)

	return output

}
