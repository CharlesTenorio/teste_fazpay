package product

import (
	"net/http"

	"github.com/fazpay/back-end/api-product/pkg/service/product"
	"github.com/go-chi/chi/v5"
)

func RegisterProductAPIHandlers(r chi.Router, service product.ProductServiceInterface) {
	r.Route("/api/v1/prd", func(r chi.Router) {
		r.Post("/product", createProduct(service))
		r.Patch("/product/", updateProduct(service))
		r.Delete("/product/", deleteProduct(service))
		r.Get("/product/", getProduct(service))
		r.Get("/products", func(w http.ResponseWriter, r *http.Request) {
			handler := getAllProducts(service)
			handler.ServeHTTP(w, r)
		})
	})
}
