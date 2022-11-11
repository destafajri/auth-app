package helper

import (
	"encoding/json"
	
	"github.com/gin-gonic/gin"
)

func ResponseJSON(w gin.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}