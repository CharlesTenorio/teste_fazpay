package main

import (
	"log"
	"net/http"

	"github.com/fazpay/back-end/api-product/internal/config"
	"github.com/fazpay/back-end/api-product/internal/config/logger"

	_ "github.com/fazpay/back-end/api-product/docs"
	hand_prd "github.com/fazpay/back-end/api-product/internal/handler/product"
	"github.com/fazpay/back-end/api-product/pkg/adapter/pgsql"
	"github.com/fazpay/back-end/api-product/pkg/server"
	service_prd "github.com/fazpay/back-end/api-product/pkg/service/product"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

var (
	VERSION = "0.1.0-dev"
	COMMIT  = "ABCDEFG-dev"
)

// @title           Fazpay API
// @version         1.0
// @description     API for Faz.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  charles.tenorio.dev@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {

	logger.Info("start Notifaction application")
	conf := config.NewConfig()
	conn_pg := pgsql.New(conf)

	prd_service := service_prd.NewProductService(conn_pg)
	r := chi.NewRouter()

	r.Get("/", healthcheck)
	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/docs/doc.json")))

	hand_prd.RegisterProductAPIHandlers(r, prd_service)
	srv := server.NewHTTPServer(r, conf)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	log.Printf("Server Run on [Port: %s], [Mode: %s], [Version: %s], [Commit: %s]", conf.PORT, conf.Mode, VERSION, COMMIT)

	select {}
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"MSG": "Server Ok", "codigo": 200}`))
}
