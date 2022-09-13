package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func ApiServer() {

	var router = mux.NewRouter().StrictSlash(true)

	router.NotFoundHandler = http.HandlerFunc(ErrorNotFound)

	var resource = router.PathPrefix("/v1").Subrouter()

	resource.Use(RequestHandler)

	resource.HandleFunc("/", ResourceIndex)

	log.Println("Listening on Port: 2409")

	err := http.ListenAndServe(":2409", router)
	log.Fatal(err)
}
