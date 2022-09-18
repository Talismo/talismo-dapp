package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func ResourceWallets(w http.ResponseWriter, r *http.Request) {

	var wallets *DefaultApiResponse = &DefaultApiResponse{
		Status:  "200",
		Message: "Looking for wallets?",
	}

	jsonResponse, _ := json.Marshal(wallets)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}

func ResourceWallet(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

	if err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(responseData))
}
