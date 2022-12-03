package Controller

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Categorias() string {

	resp, err := http.Get("https://api.chucknorris.io/jokes/categories")

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	output := string(responseData)
	fmt.Println(output)

	return output

}
