package product

import (
	"encoding/json"
	"net/http"

	"github.com/fazpay/back-end/api-product/internal/config/logger"
	"github.com/fazpay/back-end/api-product/pkg/model"
	"github.com/fazpay/back-end/api-product/pkg/service/product"
)

// getAllProducts retorna todos os produtos.
// @Summary Retorna todos os produtos
// @Description Retorna todos os produtos
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {array} model.Product
// @Failure 500 {object} handler.HttpMsg
// @Router /prd/all [get]
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

// getProduct retorna um produto pelo seu ID externo.
// @Summary Retorna um produto pelo seu ID externo
// @Description Retorna um produto com base no ID externo fornecido no cabeçalho da requisição
// @Tags producs
// @Accept json
// @Produce json
// @Param id header string true "ID externo do produto"
// @Success 200 {object} model.Product
// @Failure 400 {object} handler.HttpMsg "O campo ID é obrigatório"
// @Failure 404 {object} handler.HttpMsg "Produto não encontrado"
// @Failure 500 {object} handler.HttpMsg "Erro ao converter produto para JSON"
// @Router /prd/getbyid [get]
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

// createProduct cria um novo produto.
// @Summary Cria um novo produto
// @Description Cria um novo produto com base nos dados fornecidos no corpo da requisição
// @Tags products
// @Accept json
// @Produce json
// @Param product body model.Product true "Dados do produto a ser criado"
// @Success 200 {object} model.Product
// @Failure 400 {object} handler.HttpMsg "Campos obrigatórios não foram fornecidos"
// @Failure 500 {object} handler.HttpMsg "Erro ao criar o produto"
// @Router /prd/add [post]
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

// updateProduct atualiza um produto existente.
// @Summary Atualiza um produto existente
// @Description Atualiza um produto com base no ID externo fornecido no cabeçalho da requisição
// @Tags products
// @Accept json
// @Produce json
// @Param id header string true "ID externo do produto"
// @Param product body model.Product true "Dados do produto a ser atualizado"
// @Success 200 {object} model.Product
// @Failure 400 {object} handler.HttpMsg "O campo ID é obrigatório"
// @Failure 404 {object} handler.HttpMsg "Produto não encontrado"
// @Failure 500 {object} handler.HttpMsg "Erro ao atualizar o produto"
// @Router /prd/update/ [patch]
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

// deleteProduct exclui um produto existente.
// @Summary Exclui um produto existente
// @Description Exclui um produto com base no ID externo fornecido no cabeçalho da requisição
// @Tags products
// @Accept json
// @Produce json
// @Param id header string true "ID externo do produto"
// @Success 200 {object} handler.HttpMsg "Produto excluído com sucesso"
// @Failure 400 {object} handler.HttpMsg "O campo ID é obrigatório"
// @Failure 404 {object} handler.HttpMsg "Produto não encontrado"
// @Failure 500 {object} handler.HttpMsg "Erro ao excluir o produto"
// @Router /prd/delete/ [delete]
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
