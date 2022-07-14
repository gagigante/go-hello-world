package main

import (
	"encoding/json"
	"net/http"
)

type Task struct {
	Name string
	Done bool
}

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/task", TaskHandler)
	http.ListenAndServe(":8888", nil)
}

func Hello(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello world from my first web server"))
}

func TaskHandler(writer http.ResponseWriter, request *http.Request) {
	task := Task{
		Name: "Setup",
		Done: false,
	}

	j, _ := json.Marshal(task)
	writer.Write(j)

	// t := template.Must(template.ParseFiles("task.html"))
	// t.Execute(writer, task)
}
