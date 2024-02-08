package product

import (
	"encoding/json"
	"net/http"

	"github.com/fazpay/back-end/api-product/internal/config/logger"
	"github.com/fazpay/back-end/api-product/pkg/model"
	"github.com/fazpay/back-end/api-product/pkg/service/product"
)

func getAllProducts(service product.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		list_products := service.GetAll(r.Context())
		err := json.NewEncoder(w).Encode(list_products)
		if err != nil {
			ErroHttpMsgToConvertingResponseProductListToJson.Write(w)
			return
		}
	})
}

func getProduct(service product.ProductServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		externalID := r.Header.Get("id")

		if externalID == "" {
			ErroHttpMsgProductIdIsRequired.Write(w)
			return
		}

		product := service.GetByExternalID(r.Context(), externalID)
		if product.ID == 0 {
			ErroHttpMsgProductNotFound.Write(w)
			return
		}

		err := json.NewEncoder(w).Encode(product)
		if err != nil {
			ErroHttpMsgToParseResponseProductToJson.Write(w)
			return
		}
	}
}

func createProduct(service product.ProductServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		product := model.Product{}

		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			ErroHttpMsgToParseRequestProductToJson.Write(w)
			return
		}

		prd, err := model.NewProduct(&product)

		if err != nil {
			ErroHttpMsgToParseRequestProductToJson.Write(w)
			return
		}

		logger.Info(prd.Name)
		if prd.Name == " " || prd.Name == "" {
			ErroHttpMsgProductNameIsRequired.Write(w)
			return
		}

		if prd.Quantity <= 0 {
			ErroHttpMsgProductQuantityIsRequired.Write(w)
			return
		}

		if prd.Price <= 0.0 {
			ErroHttpMsgProductPriceIsRequired.Write(w)
			return
		}
		product = *prd

		external_id := service.Create(r.Context(), &product)
		if external_id == "" {
			ErroHttpMsgToInsertProduct.Write(w)
			return
		}

		product = *service.GetByExternalID(r.Context(), external_id)

		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			ErroHttpMsgToParseResponseProductToJson.Write(w)
			return
		}
	}
}
func updateProduct(service product.ProductServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		externalID := r.Header.Get("id")

		if externalID == "" {
			ErroHttpMsgProductIdIsRequired.Write(w)
			return
		}

		request_to_update_product := model.Product{}

		err := json.NewDecoder(r.Body).Decode(&request_to_update_product)
		if err != nil {
			ErroHttpMsgToParseRequestProductToJson.Write(w)
			return
		}

		old_product := service.GetByExternalID(r.Context(), externalID)

		if old_product.ID == 0 {
			ErroHttpMsgProductNotFound.Write(w)
			return
		}

		rows_affected := service.Update(r.Context(), old_product.ID, &request_to_update_product)
		if rows_affected == 0 {
			ErroHttpMsgToUpdateProduct.Write(w)
			return
		}

		request_to_update_product.ExternalID = old_product.ExternalID

		err = json.NewEncoder(w).Encode(request_to_update_product)
		if err != nil {
			ErroHttpMsgToParseResponseProductToJson.Write(w)
			return
		}
	}
}
func deleteProduct(service product.ProductServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		externalID := r.Header.Get("id")

		if externalID == "" {
			ErroHttpMsgProductIdIsRequired.Write(w)
			return
		}

		old_product := service.GetByExternalID(r.Context(), externalID)

		if old_product.ID == 0 {
			ErroHttpMsgProductNotFound.Write(w)
			return
		}

		rows_affected := service.Delete(r.Context(), old_product.ID)
		if rows_affected == 0 {
			ErroHttpMsgToDeleteProduct.Write(w)
			return
		}
		SuccessHttpMsgToDeleteProduct.Write(w)
	}
}
