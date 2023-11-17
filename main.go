package main

import (
	"log"
	"net/http"
	"strconv"
)

func main() {
	port := 8080

	fs := http.FileServer(http.Dir("./web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/tasks", getTasks)
	http.HandleFunc("/add-task", addTask)
	http.HandleFunc("/toggle-status", toggleTaskStatus)
	http.HandleFunc("/remove", removeTask)
	http.HandleFunc("/", indexHandler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}
