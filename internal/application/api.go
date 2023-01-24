package application

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"dgb/meter.readings.bot/internal/configuration"

	"github.com/gorilla/mux"
)

func HandleRequests(conf configuration.Configuration) chan int {

	r := make(chan int)

	myRouter := mux.NewRouter().StrictSlash(true)
	subRoute := myRouter.PathPrefix("/api").Subrouter()

	subRoute.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode("OK")
	}).Methods(http.MethodGet)

	go func() {
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", conf.HTTP_PORT), subRoute))
	}()

	return r
}
