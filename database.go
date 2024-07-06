package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func getTodos() []Todo {
	projectUrl := getEnvVariable("PROJECT_URL")
	apiKey := getEnvVariable("API_KEY")

	client := &http.Client{}
	url := projectUrl + "/todos?order=date.desc,created_at.desc&date=eq." + time.Now().Format("2006-01-02")
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("apikey", apiKey)

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()

	if resp.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
	}

	if err != nil {
		log.Fatal(err)
	}

	var todos []Todo
	err = json.Unmarshal(body, &todos)

	if err != nil {
		log.Fatal(err)
	}

	return todos
}

func createTodo(description string) error {
	projectUrl := getEnvVariable("PROJECT_URL")
	apiKey := getEnvVariable("API_KEY")
	client := &http.Client{}

	url := projectUrl + "/todos"
	payload := []byte(fmt.Sprintf(`{"description":"%s"}`, description))
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))

	req.Header.Set("apikey", apiKey)
	req.Header.Set("Content-Type", "application/json")

	_, err := client.Do(req)

	if err != nil {
		return err
	}

	return nil
}
