package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fazpay/back-end/api-product/internal/config"
	"github.com/go-chi/chi/v5"
)

func NewHTTPServer(r chi.Router, conf *config.Config) *http.Server {
	srv := &http.Server{
		ReadTimeout:  10 * time.Second, // Wait for 10 seconds for a request to be fully read
		WriteTimeout: 10 * time.Second, // Respond within 10 seconds
		Addr:         ":" + conf.PORT,
		Handler:      r,
		ErrorLog:     log.New(os.Stderr, "logger: ", log.Lshortfile),
	}

	return srv
}
