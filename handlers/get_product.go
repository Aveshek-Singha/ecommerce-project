package handler

import (
	"ecommerice-project/database"
	"ecommerice-project/util"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	util.SendData(w, database.ProductList, 200)
}
