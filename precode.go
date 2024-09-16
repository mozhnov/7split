package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Task ...
type Task struct {
	ID           string   `json:"id"`
	Description  string   `json:"description"`
	Note         string   `json:"note"`
	Applications []string `json:"applications"`
}

var tasks = map[string]Task{
	"1": {
		ID:          "1",
		Description: "Сделать финальное задание темы REST API",
		Note:        "Если сегодня сделаю, то завтра будет свободный день. Ура!",
		Applications: []string{
			"VS Code",
			"Terminal",
			"git",
		},
	},
	"2": {
		ID:          "2",
		Description: "Протестировать финальное задание с помощью Postmen",
		Note:        "Лучше это делать в процессе разработки, каждый раз, когда запускаешь сервер и проверяешь хендлер",
		Applications: []string{
			"VS Code",
			"Terminal",
			"git",
			"Postman",
		},
	},
}

// Ниже напишите обработчики для каждого эндпоинта
func getTasks(w http.ResponseWriter, r *http.Request) {
	// сериализуем данные из мапы tasks
	err := json.NewEncoder(w).Encode(tasks)
	//проверяем на ошибки
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// возвращаем ответ клиенту
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
func getTaskId(w http.ResponseWriter, r *http.Request) {
	// получаем id
	id := r.URL.Query().Get("id")
	//создаем переменную
	task, ok := tasks[id]
	//проверяем на ошибки
	if !ok {
		http.Error(w, "Задача не найдена", http.StatusBadRequest)
		return
	}
	// сериализуем данные из переменной task по id
	json.NewEncoder(w).Encode(task)
	// возвращаем ответ клиенту
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
func postTask(w http.ResponseWriter, r *http.Request) {
	//создаем переменную
	var task Task
	// десериализуем данные из Request и записываем в указатель на переменную task
	err := json.NewDecoder(r.Body).Decode(&task)
	//проверяем на ошибки
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// добавляем новую задачу в мапу
	tasks[task.ID] = task
	// возвращаем ответ клиенту
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
func deleteTask(w http.ResponseWriter, r *http.Request) {
	// получаем id
	id := r.URL.Query().Get("id")
	//создаем переменную
	task, ok := tasks[id]
	//проверяем на ошибки
	if !ok {
		http.Error(w, "Задача не найдена", http.StatusBadRequest)
		return
	}
	// удаляем задачу из мапы tasks
	delete(tasks, task.ID)
	// возвращаем ответ клиенту
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func main() {
	r := chi.NewRouter()

	// здесь регистрируйте ваши обработчики
	r.Get("/tasks", getTasks)
	r.Get("/tasks/{id}", getTaskId)
	r.Post("/tasks", postTask)
	r.Delete("/tasks/{id}", deleteTask)

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s", err.Error())
		return
	}
}
