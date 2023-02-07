package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	baseUrl = "http://192.168.0.173:8080"
)

type Data struct {
	Content string `json:"content"`
}

func main() {

	data := Data{
		Content: os.Args[1],
	}

	_, err := do(context.Background(), data)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()
}

func do(ctx context.Context, data Data) (*http.Response, error) {

	buf := &bytes.Buffer{}

	if err := json.NewEncoder(buf).Encode(&data); err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("%s/clip", baseUrl), buf)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Source-App", "client-go")

	client := http.DefaultClient

	return client.Do(req)
}
