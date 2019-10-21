package main

type Employee struct {
	Employee_ID		string	`form:"employee_id" json:"employee_id"`
	Full_Name		string 	`form:"full_name" json:"full_name"`
	Department		string	`form:"department" json:"department"`
	Designation		string	`form:"designation" json:"designation"`
	Status			string	`form:"status" json:"status"`
}

type Response struct {
	Status	int		`json:"status"`
	Message string	`json:"message"`
	Data 	[]Employee
}