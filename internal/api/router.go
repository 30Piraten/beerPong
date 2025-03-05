package api

import "github.com/gorilla/mux"

func Router() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/throw", throwBallHandler).Methods("POST")
	r.HandleFunc("/cup/{cup_id}", cupHandler).Methods("GET")

	return r
}
