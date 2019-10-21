package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/employee", getEmployee).Methods("GET")
	router.HandleFunc("/employes", selectEmployee).Methods("GET")
	router.HandleFunc("/employee", insertEmployee).Methods("POST")
	router.HandleFunc("/employee", updateEmployee).Methods("PUT")
	router.HandleFunc("/employee", deleteEmployee).Methods("DELETE")
	http.Handle("/", router)
	fmt.Println("Connected to port 8181")
	log.Fatal(http.ListenAndServe(":8181", router))
}

//type Employee struct {
//	Employee_ID		string	`json"employee_id"`
//	Full_Name		string 	`json"full_name"`
//	Department		string	`json"departmen"`
//	Designation		string	`json"designation"`
//	Status			string	`json"status"`
//}
//
//var employee []Employee
//
//func getEmployee(w http.ResponseWriter, r *http.Request){
//	w.Header().Set("Content-Type", "application/json")
//	json.NewEncoder(w).Encode(employee)
//
//	db := con
//}

//func main() {
//	// Connect to the "company_db" database.
//	db, err := sql.Open("postgres", "postgresql://root@localhost:26257/company_db?sslmode=disable")
//	if err != nil {
//		log.Fatal("error connecting to the database: ", err)
//	}
//
//
//	r := mux.NewRouter()
//	r.HandleFunc("/api/employee", getEmployee).Methods("GET")
//	log.Fatal(http.ListenAndServe(":8080", r))
//
//	Insert a row into the "tbl_employee" table.
//	if _, err := db.Exec(
//		`INSERT INTO employee (employee_id,full_name, department, designation, status)
//		VALUES ('0000000001','David', 'IT', 'Programmer', 'active');`); err != nil {
//		log.Fatal(err)
//	}
//
//	Select Statement.
//	rows, err := db.Query("select employee_id, full_name FROM employee;")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer rows.Close()
//	for rows.Next() {
//		var employeeID int64
//		var fullName string
//		if err := rows.Scan(&employeeID, &fullName); err != nil {
//			log.Fatal(err)
//		}
//		fmt.Printf("Employee Id : %d \t Employee Name : %s\n", employeeID, fullName)
//	}
//}