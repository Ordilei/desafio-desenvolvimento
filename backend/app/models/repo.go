package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func listAllRepo(w http.ResponseWriter, r *http.Request) {

	cliente := &http.Client{
		Timeout: time.Second * 30,
	}
	page := 1
	resposta, err := cliente.Get("https://api.github.com/users/Ordilei")

	for page != 0 {

		if err != nil {
			fmt.Println("Erro =>", err.Error())
			return
		}
		defer resposta.Body.Close()

		page++
	}

	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("Erro: ", err.Error())
			return
		}
		fmt.Fprintf(w, "%s", corpo)
	}
}

func main() {

	http.HandleFunc("/list-all", listAllRepo)
	log.Println("server up")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
