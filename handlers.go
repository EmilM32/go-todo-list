package main

import (
	"html/template"
	"net/http"

	"github.com/google/uuid"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/index.html")
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/task-form.gohtml", "web/tasks-page.gohtml")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Tasks []Task
	}{
		Tasks: tasks,
	}

	tmpl.ExecuteTemplate(w, "tasks-page", data)
}

func addTask(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	title := r.FormValue("title")
	task := Task{ID: uuid.New().String(), Title: title, Done: false}
	tasks = append(tasks, task)

	getTasks(w, r)
}

func toggleTaskStatus(w http.ResponseWriter, r *http.Request) {
	taskID := r.FormValue("taskID")

	for i, task := range tasks {
		if task.ID == taskID {
			tasks[i].Done = !tasks[i].Done
			break
		}
	}

	getTasks(w, r)
}

func removeTask(w http.ResponseWriter, r *http.Request) {
	taskID := r.FormValue("taskID")

	for i, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}

	getTasks(w, r)
}
