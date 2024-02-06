package controllers

import (
	"app/controllers/tasks"

	"github.com/gorilla/mux"
)

func Serve(r *mux.Router) {
	r.HandleFunc("/tasks", tasks.GetAll )			.Methods("GET")
	r.HandleFunc("/tasks/id/{id}", tasks.GetOne )	.Methods("GET")
	r.HandleFunc("/tasks", tasks.CreateOne )		.Methods("POST")
	r.HandleFunc("/tasks/{id}", tasks.UpdateOne )	.Methods("PUT")
	r.HandleFunc("/tasks/{id}", tasks.DeleteOne )	.Methods("DELETE")
}