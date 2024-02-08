package product

import (
	"net/http"

	"github.com/fazpay/back-end/api-product/pkg/service/product"
	"github.com/go-chi/chi/v5"
)

func RegisterProductAPIHandlers(r chi.Router, service product.ProductServiceInterface) {
	r.Route("/api/v1/prd", func(r chi.Router) {
		r.Post("/add", createProduct(service))
		r.Patch("/update/", updateProduct(service))
		r.Delete("/delete/", deleteProduct(service))
		r.Get("/getbyid/", getProduct(service))
		r.Get("/all", func(w http.ResponseWriter, r *http.Request) {
			handler := getAllProducts(service)
			handler.ServeHTTP(w, r)
		})
	})
}
