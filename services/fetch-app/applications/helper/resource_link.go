package helper

import (
	"encoding/json"
	"fmt"
	"io"
	_ "log"
	"net/http"
	"os"
)

func Resource() []ResponseFetch{
	fmt.Println("Calling API...")
	
	client := &http.Client{}
	req, err := http.NewRequest("GET", os.Getenv("URL"), nil)
	
	if err != nil {
		fmt.Print(err.Error())
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	
	defer resp.Body.Close()
	
	bodyBytes, err := io.ReadAll(resp.Body)
	
	if err != nil {
		fmt.Print(err.Error())
	}
	
	getFetch := make([]ResponseFetch, 0)

	json.Unmarshal(bodyBytes, &getFetch)

	return getFetch
}