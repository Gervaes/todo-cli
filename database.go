package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

func getTodos(getAllTodos bool) []Todo {
	projectUrl := getEnvVariable("TODOS_PROJECT_URL")
	apiKey := getEnvVariable("TODOS_API_KEY")

	client := &http.Client{}
	url := projectUrl + "/todos?order=date.desc,created_at.desc"
	if !getAllTodos {
		url += "&date=eq." + time.Now().Format("2006-01-02")
	}
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

func getTodo(id int) (Todo, error) {
	projectUrl := getEnvVariable("TODOS_PROJECT_URL")
	apiKey := getEnvVariable("TODOS_API_KEY")

	client := &http.Client{}
	url := projectUrl + "/todos?id=eq." + strconv.Itoa(id)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("apikey", apiKey)

	resp, err := client.Do(req)

	if err != nil {
		return Todo{}, err
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()

	if resp.StatusCode > 299 {
		return Todo{}, *new(error)
	}

	if err != nil {
		return Todo{}, err
	}

	var todos []Todo
	err = json.Unmarshal(body, &todos)

	if err != nil {
		return Todo{}, err
	}

	if len(todos) < 1 || len(todos) > 1 {
		return Todo{}, *new(error)
	}

	todo := todos[0]

	return todo, nil
}

func createTodo(description string) error {
	projectUrl := getEnvVariable("TODOS_PROJECT_URL")
	apiKey := getEnvVariable("TODOS_API_KEY")
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

func updateTodo(updatedTodo Todo) error {
	todo, err := json.Marshal(updatedTodo)

	if err != nil {
		return err
	}

	projectUrl := getEnvVariable("TODOS_PROJECT_URL")
	apiKey := getEnvVariable("TODOS_API_KEY")
	client := &http.Client{}

	url := projectUrl + "/todos?id=eq." + strconv.FormatInt(int64(updatedTodo.Id), 10)
	req, _ := http.NewRequest("PATCH", url, bytes.NewBuffer(todo))

	req.Header.Set("apikey", apiKey)
	req.Header.Set("Content-Type", "application/json")

	_, err = client.Do(req)

	if err != nil {
		return err
	}

	return nil
}

func deleteTodo(id int) error {
	projectUrl := getEnvVariable("TODOS_PROJECT_URL")
	apiKey := getEnvVariable("TODOS_API_KEY")
	client := &http.Client{}

	url := projectUrl + "/todos?id=eq." + strconv.Itoa(id)
	req, _ := http.NewRequest("DELETE", url, nil)
	req.Header.Set("apikey", apiKey)

	_, err := client.Do(req)

	if err != nil {
		return err
	}

	return nil
}
