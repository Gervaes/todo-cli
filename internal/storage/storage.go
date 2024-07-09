package storage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	"todo-cli/internal/models"
)

type Storage struct {
	ProjectUrl string
	ApiKey     string
	HttpClient *http.Client
}

func NewStorage() Storage {
	store := Storage{
		ProjectUrl: os.Getenv("TODOS_PROJECT_URL"),
		ApiKey:     os.Getenv("TODOS_API_KEY"),
		HttpClient: &http.Client{},
	}

	return store
}

func (s *Storage) GetTodos(getAllTodos bool) []models.Todo {
	url := s.ProjectUrl + "/todos?order=date.desc,created_at.desc"
	if !getAllTodos {
		url += "&date=eq." + time.Now().UTC().Format("2006-01-02")
	}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("apikey", s.ApiKey)

	resp, err := s.HttpClient.Do(req)

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

	var todos []models.Todo
	err = json.Unmarshal(body, &todos)

	if err != nil {
		log.Fatal(err)
	}

	return todos
}

func (s *Storage) GetTodo(id int) (models.Todo, error) {
	url := s.ProjectUrl + "/todos?id=eq." + strconv.Itoa(id)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("apikey", s.ApiKey)

	resp, err := s.HttpClient.Do(req)

	if err != nil {
		return models.Todo{}, err
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()

	if resp.StatusCode > 299 {
		return models.Todo{}, *new(error)
	}

	if err != nil {
		return models.Todo{}, err
	}

	var todos []models.Todo
	err = json.Unmarshal(body, &todos)

	if err != nil {
		return models.Todo{}, err
	}

	if len(todos) < 1 || len(todos) > 1 {
		return models.Todo{}, *new(error)
	}

	todo := todos[0]

	return todo, nil
}

func (s *Storage) CreateTodo(description string) error {
	url := s.ProjectUrl + "/todos"
	payload := []byte(fmt.Sprintf(`{"description":"%s"}`, description))
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))

	req.Header.Set("apikey", s.ApiKey)
	req.Header.Set("Content-Type", "application/json")

	_, err := s.HttpClient.Do(req)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) UpdateTodo(updatedTodo models.Todo) error {
	todo, err := json.Marshal(updatedTodo)

	if err != nil {
		return err
	}

	url := s.ProjectUrl + "/todos?id=eq." + strconv.FormatInt(int64(updatedTodo.Id), 10)
	req, _ := http.NewRequest("PATCH", url, bytes.NewBuffer(todo))

	req.Header.Set("apikey", s.ApiKey)
	req.Header.Set("Content-Type", "application/json")

	_, err = s.HttpClient.Do(req)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) DeleteTodo(id int) error {
	url := s.ProjectUrl + "/todos?id=eq." + strconv.Itoa(id)
	req, _ := http.NewRequest("DELETE", url, nil)
	req.Header.Set("apikey", s.ApiKey)

	_, err := s.HttpClient.Do(req)

	if err != nil {
		return err
	}

	return nil
}
