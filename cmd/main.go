package main

import (
	"encoding/json"
	"net/http"
)

type Student struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var students []Student

// Get all students
func getStudents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(students)
}

// Add student
func addStudent(w http.ResponseWriter, r *http.Request) {
	var s Student
	json.NewDecoder(r.Body).Decode(&s)
	students = append(students, s)
	json.NewEncoder(w).Encode(s)
}

func main() {
	http.HandleFunc("/students", getStudents)
	http.HandleFunc("/add", addStudent)

	http.ListenAndServe(":8080", nil)
}