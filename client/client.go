package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Query struct {
	Message string `json:"message"`
}

type Response struct {
	Data struct {
		Message string `json:"message"`
	} `json:"data"`
}

func main() {
	query := `{"query": "{ message(id: \"1\") }"}`

	req, err := http.NewRequest("POST", "http://localhost:8080/graphql", bytes.NewBuffer([]byte(query)))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var result Response
	json.NewDecoder(resp.Body).Decode(&result)

	fmt.Println("Response from server:", result.Data.Message)
}
