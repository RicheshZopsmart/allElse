package Net_rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Employee struct {
	Id   int
	Name string
}

var e_obj []Employee

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	// ids := r.URL.Query().Get("id")
	// int_id, _ := strconv.Atoi(ids)
	name := r.URL.Query().Get("name")
	fmt.Println("Id : ", len(e_obj)+1)
	fmt.Println("Name : ", name)
	e := Employee{
		Id:   len(e_obj) + 1,
		Name: name,
	}
	e_obj = append(e_obj, e)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"Success"}`))
}

func GetEmployee(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	fmt.Println("Get id : ", id)
	w.Header().Set("Content-Type", "application/json")

	var fl bool = false

	if len(id) == 0 {
		for _, val := range e_obj {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(val)
			fmt.Println("Name : ", val.Name)
			fl = true
		}
	}

	int_id, _ := strconv.Atoi(id)
	fmt.Println("id : ", int_id)

	for _, val := range e_obj {
		if val.Id == int_id {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(val)
			fmt.Println("Name : ", val.Name)
			fl = true
		}
	}

	if !fl {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"Error": "Invalid_Id"}`))
	}
}
