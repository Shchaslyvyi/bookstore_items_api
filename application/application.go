package application

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/shchaslyvyi/bookstore_items_api/clients/elasticsearch"
	"github.com/shchaslyvyi/bookstore_users-api/logger"
)

var (
	router = mux.NewRouter()
)

// StartApplication func starts the application
func StartApplication() {
	elasticsearch.Init()
	mapUrls()

	srv := &http.Server{
		Addr: ":8081",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      router,
	}

	logger.Info("about to start the application...")
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
