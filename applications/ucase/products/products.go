package products

import (
	"net/http"

	"github.com/destafajri/auth-app/applications/helper"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := []map[string]interface{}{
		{
			"id":           1,
			"nama_product": "Iphone 14",
			"stok":         1000,
		},
		{
			"id":           2,
			"nama_product": "Ipad Air 20",
			"stok":         10000,
		},
		{
			"id":           3,
			"nama_product": "Ipad Air 500",
			"stok":         500,
		},
	}

	helper.ResponseJSON(w, http.StatusOK, data)
}