package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func executeRequest() {
	client := &http.Client{Timeout: 10 * time.Second}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.github.com", nil)
	if err != nil {
		panic(fmt.Errorf("build request: %w", err))
	}
	
	req.Header.Set("User-Agent", "golang-http-client-example")

	resp, err := client.Do(req)
	if err != nil {
		panic(fmt.Errorf("send request: %w", err))
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(fmt.Errorf("read body: %w", err))
	}

	fmt.Printf("Status: %s\n", resp.Status)
	if len(body) > 400 {
		fmt.Printf("Body (first 400 bytes):\n%s\n", string(body[:400]))
		return
	}
	fmt.Printf("Body:\n%s\n", string(body))
}


