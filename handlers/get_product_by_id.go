package handler

import (
	"ecommerice-project/database"
	"ecommerice-project/util"
	"net/http"
	"strconv"
)

func GetProductById(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("productId")

	pId, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please give me a valid product id", 400)
		return
	}

	for _, product := range database.ProductList {
		if product.ID == pId {
			util.SendData(w, product, 200)
			return
		}
	}

	util.SendData(w, "Data pai nai", 404)

}
