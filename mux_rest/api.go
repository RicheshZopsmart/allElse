package mux_rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Employee struct {
	Id    int
	Name  string
	Email string
	Role  string
}

var e_obj []Employee

func GetEmployee(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, ok := vars["id"]
	// id := r.URL.Query().Get("id")

	w.Header().Set("Content-Type", "application/json")

	fmt.Println("Id: ", id)

	var full_list bool = false

	if !ok {
		full_list = true
	}

	int_id, err := strconv.Atoi(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"Error": "Invalid_Id"}`))
	}

	var found bool = false

	for _, val := range e_obj {
		if full_list || val.Id == int_id {
			found = true
			json.NewEncoder(w).Encode(val)
		}
	}

	if !found {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"Error": "No_User_Found"}`))
	}

}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {

	// vars := mux.Vars(r)
	// name := r.URL.Query().Get("name")
	// name := vars["name"]

	var e Employee
	_ = json.NewDecoder(r.Body).Decode(&e)
	id := len(e_obj) + 1
	e.Id = id
	e_obj = append(e_obj, e)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(e)

}

func GetEmployees(w http.ResponseWriter, r *http.Request) {

	if len(e_obj) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"Error": "No_Data_Found"}`))
	}

	for _, val := range e_obj {

		json.NewEncoder(w).Encode(val)

	}
}

func UpdateEmployees(w http.ResponseWriter, r *http.Request) {

	var e Employee

	_ = json.NewDecoder(r.Body).Decode(&e)

	if e.Id == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"Error": "No_Data_Found"}`))
	}

	var tobeupdated *Employee

	for idx, val := range e_obj {
		if val.Id == e.Id {
			tobeupdated = &e_obj[idx]
			break
		}
	}

	tobeupdated.Id = e.Id
	tobeupdated.Name = e.Name
	tobeupdated.Email = e.Email
	tobeupdated.Role = e.Role

	w.Write([]byte(`{"Success"}`))

}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	int_id, _ := strconv.Atoi(id)

	var fl bool = false

	for idx, val := range e_obj {
		if val.Id == int_id {
			e_obj = append(e_obj[:idx], e_obj[idx+1:]...)
			fl = true
		}

	}
	if fl {
		w.Write([]byte(`{"Success"}`))
	} else {
		w.Write([]byte(`{"Error":"User_Not_Found"}`))
	}
}
