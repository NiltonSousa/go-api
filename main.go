package main

import (
	"fmt"
	"net/http"

	"github.com/NiltonSousa/goapi/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)

	// Creating chi variable to use on handlers to customize routers
	var r *chi.Mux = chi.NewRouter()

	// Calling handler
	handlers.Handler(r)

	fmt.Println("Starting GO API service...")

	// Creating server
	err := http.ListenAndServe("localhost:8000", r)

	if err != nil {
		log.Error(err)
	}
}
