package main

import (
	"log"
	"net/http"

	config "CMS_NEW/Database"
	routers "CMS_NEW/Routers"
)

func main() {

	// Load the configuration stuffs.
	config.LoadAppConfig()

	// Initiate the database information.
	config.CreateDBSession()

	// Add the neccessary indexes.
	config.AddIndexes()

	// Created the routes of this application/
	mux := http.NewServeMux()
	mux = routers.SetCandidateRoutes(mux)

	log.Println("Listening...")

	// Listen the server.
	http.ListenAndServe(":9090", mux)
}
