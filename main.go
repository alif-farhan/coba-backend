package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/osmanzz/workshop-backend/service"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/select_all_data", service.GetAllData).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})

	handler := c.Handler(router)

	log.Println("======== successfull listening to :8022")
	http.ListenAndServe(":8022", handler)
}
