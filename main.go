package main

import (
	// "database/sql"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	crud "github.com/RicheshZopsmart/Learning-go/CRUDSql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Employee struct {
	Id    int
	Name  string
	Email string
	Role  string
}

var db *sql.DB
var e_obj []Employee

func GetEmployee(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, ok := vars["id"]
	int_id, err := strconv.Atoi(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"Error": "Invalid_Id"}`))
	}

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"Error": "No_USER_FOUND"}`))
	}

	data, dberr := crud.GetById(db, int_id)

	if dberr != nil {
		w.Write([]byte(`{"Error": "Invalid_Id"}`))
	}
	w.Header().Set("Content-Type", "application/json")

	fmt.Println("Id: ", id)

	json.NewEncoder(w).Encode(data)

}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {

	var e Employee
	_ = json.NewDecoder(r.Body).Decode(&e)

	w.Header().Set("Content-Type", "application/json")
	crud.InsertData("Employee_Details", db, e.Name, e.Email, e.Role)

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
	crud.UpdateById(db, e.Id, e.Name, e.Email, e.Role)

	w.Write([]byte(`{"Success"}`))

}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	int_id, _ := strconv.Atoi(id)
	fl := crud.DeleteById(db, int_id)

	if fl == nil {
		w.Write([]byte(`{"Success"}`))
	} else {
		w.Write([]byte(`{"Error":"User_Not_Found"}`))
	}
}

func main() {

	// Gorilla/mux apis

	db = crud.DbConn("Employee_Db")
	err := db.Ping()
	if err != nil {
		log.Fatal("DB connection error ")
	}

	r := mux.NewRouter()
	r.HandleFunc("/emp/", CreateEmployee).Methods("POST")
	r.HandleFunc("/emp/{id}", GetEmployee).Methods("GET")
	r.HandleFunc("/emp", GetEmployees).Methods("GET")
	r.HandleFunc("/emp", UpdateEmployees).Methods("PATCH")
	r.HandleFunc("/emp/{id}", DeleteEmployee).Methods("DELETE")
	http.ListenAndServe(":8090", r)

}
