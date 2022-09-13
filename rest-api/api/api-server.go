package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RequestHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func ApiServer() {

	var router = mux.NewRouter().StrictSlash(true)

	router.NotFoundHandler = http.HandlerFunc(ErrorNotFound)

	var resource = router.PathPrefix("/v1").Subrouter()

	resource.Use(RequestHandler)

	resource.HandleFunc("/", ResourceIndex)

	log.Println("Listening on Port: 2409")

	http.ListenAndServe(":2409", router)
}
