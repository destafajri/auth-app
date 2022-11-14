package helper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func ConvertCurrency(s string) (ResponseCurrency, error){
  var getCurrency ResponseCurrency

  urlBased := "https://api.apilayer.com/fixer/convert?to=USD&from=IDR&amount="
  url := urlBased+s

  client := &http.Client {}
  req, err := http.NewRequest("GET", url, nil)

  req.Header.Set("apikey", os.Getenv("API_KEY"))

  if err != nil {
    fmt.Println(err)
  }
  res, err := client.Do(req)
  if err != nil {
    return getCurrency, err
  }

  if res.Body != nil {
    defer res.Body.Close()
  }

  body, err := io.ReadAll(res.Body)
  if err != nil {
    return getCurrency, err
  }

  fmt.Println(string(body))

  err = json.Unmarshal(body, &getCurrency)
  if err != nil {
    return getCurrency, err
  }

  return getCurrency, nil
}
