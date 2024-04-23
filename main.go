package main

import (
	"encoding/json"
	"net/http"

	"github.com/marianaandrxde/first-crud-go/domain"
)

func main() {
	http.HandleFunc("/person/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			// criação de pessoa
			var person domain.Person
			err := json.NewDecoder(r.Body).Decode(&person)
			if err != nil {
				http.Error(w, "Error trying to create person", http.StatusBadRequest)
				return
			}

			if person.ID <= 0 {
				http.Error(w, "Error trying to create person. ID showld be a positive integer", http.StatusBadRequest)
				return
			}

			w.WriteHeader(http.StatusCreated)
			return
		}
		http.Error(w, "Not Implemented", http.StatusInternalServerError)
	})

	http.ListenAndServe(":8080", nil)
}
