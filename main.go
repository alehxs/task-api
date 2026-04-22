package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Task struct {
	ID int	`json:"id"`
	Title string `json:"title"`
	Completed bool	`json:"completed"`
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/tasks/")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	for _, task := range tasks {
		if task.ID == id{
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
			return
		}
	}

	http.NotFound(w, r)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	json.NewDecoder(r.Body).Decode(&task)

	task.ID = nextID
	nextID++

	tasks = append(tasks, task)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	var newTasks[] Task
	idStr := strings.TrimPrefix(r.URL.Path, "/tasks/")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	for _, task := range tasks {
		if task.ID != id {
			newTasks = append(newTasks, task)
		}
	}
	tasks = newTasks
	w.WriteHeader(http.StatusNoContent)
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id, err := strconv.Atoi(idStr)
	
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	
	var updatedTask Task
	json.NewDecoder(r.Body).Decode(&updatedTask)

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Title = updatedTask.Title
			tasks[i].Completed = updatedTask.Completed

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(tasks[i])
			return
		}
	}
	http.NotFound(w, r)
}

var tasks = []Task{}
var nextID = 3

func main() {
	tasks = []Task{
		{ID: 1, Title: "Eat breakfast", Completed: false},
		{ID: 2, Title: "Take a walk", Completed: true},
	}
	
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getTasks(w, r)
		case http.MethodPost:
			createTask(w, r)
		}
	})

	http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request){
		switch r.Method {
		case http.MethodGet:
			getTask(w, r)
		case http.MethodDelete:
			deleteTask(w, r)
		case http.MethodPut:
			updateTask(w, r)
		}
	})
	http.ListenAndServe(":8000", nil)
}
