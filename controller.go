package main

import (
	"encoding/json"
	"log"
	"net/http"
)

var employee Employee
var arr_employee []Employee
var arr_semploye []Employee
var response Response

func getEmployee (w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	rows, err := db.Query("Select employee_id, full_name, department, designation, status from employee")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&employee.Employee_ID, &employee.Full_Name, &employee.Department, &employee.Designation, &employee.Status);
			err != nil {
			log.Fatal(err.Error())
		} else {
			arr_employee = append(arr_employee, employee)
		}
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = arr_employee

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func selectEmployee (w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	employee_id := r.FormValue("employee_id")



	rows, err := db.Query("Select employee_id, full_name, department, designation, status from employee where employee_id = $1",
		employee_id,
	)


	for rows.Next() {
		if err := rows.Scan(&employee.Employee_ID, &employee.Full_Name, &employee.Department, &employee.Designation, &employee.Status);
			err != nil {
			log.Fatal(err.Error())
		} else {
			arr_semploye = append(arr_semploye, employee)
		}
	}

	if err != nil {
		log.Print(err)
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = arr_semploye
	arr_semploye = append(arr_semploye[1:])

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func insertEmployee (w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	employee_id := r.FormValue("employee_id")
	full_name := r.FormValue("full_name")
	department := r.FormValue("department")
	designation := r.FormValue("designation")
	status := r.FormValue("status")

	_, err = db.Exec("INSERT INTO employee (employee_id, full_name, department, designation, status) values ($1, $2, $3, $4, $5)",
		employee_id, full_name, department, designation, status,
	)

	if err != nil {
		log.Print(err)
	}

	response.Status = 2
	response.Message = "Success"
	log.Print("Insert data to database")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func updateEmployee (w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	employee_id := r.FormValue("employee_id")
	full_name := r.FormValue("full_name")
	department := r.FormValue("department")
	designation := r.FormValue("designation")

	_, err = db.Exec("UPDATE employee set full_name = $1, department = $2, designation = $3  where employee_id = $4",
		full_name, department, designation, employee_id,
	)

	if err != nil {
		log.Print(err)
	}

	response.Status = 3
	response.Message = "Update Success"
	log.Print("Update data to database")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func deleteEmployee (w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	employee_id := r.FormValue("employee_id")
	status := r.FormValue("status")

	_, err = db.Exec("UPDATE employee set status = $1  where employee_id = $2",
		status, employee_id,
	)

	if err != nil {
		log.Print(err)
	}

	response.Status = 3
	response.Message = "Delete Success"
	log.Print("Delete data from database")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}