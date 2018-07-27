package main

import "time"

//Todo is a struct to store name, completed and due
type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

//Todos is a Todo slice
type Todos []Todo
