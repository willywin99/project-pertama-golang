package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/person", personHandler)

	http.ListenAndServe("localhost:8082", nil)
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Response struct {
	Success bool   `json:"success"`
	Data    any    `json:"data"`
	Error   string `json:"error"`
}

func personHandler(awdad http.ResponseWriter, wqwq *http.Request) {

	switch wqwq.Method {
	case http.MethodGet:
		persons := make([]Person, 0)

		persons = append(persons, Person{
			Name: "Budi",
			Age:  10,
		})

		persons = append(persons, Person{
			Name: "Ani",
			Age:  20,
		})

		var r Response = Response{
			Success: true,
			Data:    persons,
		}

		jsonResponse(r, http.StatusOK, awdad)
	default:
		var r Response = Response{
			Success: false,
			Error:   "not found!",
		}

		jsonResponse(r, http.StatusNotFound, awdad)
	}
}

func jsonResponse(r Response, httpCode int, w http.ResponseWriter) {
	result, err := json.Marshal(r)
	if err != nil {
		fmt.Println("error marshalling", err)
		http.Error(w, fmt.Sprintf("error json encoding %s", err), http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	w.Write(result)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ini dari fmt.Println")
	w.Write([]byte("hello world"))
}
