package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Call API https://dummy.restapiexample.com/api/v1/employees

type Employee struct {
	ID             int    `json:"id"`
	EmployeeName   string `json:"employee_name"`
	EmployeeSalary int    `json:"employee_salary"`
	EmployeeAge    int    `json:"employee_age"`
	ProfileImage   string `json:"profile_image"`
}

type ResponseFromServer struct {
	Status string     `json:"status"`
	Data   []Employee `json:"data"`
}

func part1() {
	const apiUrl = "https://dummy.restapiexample.com/api/v1/employees"
	// Call API
	response, err := http.Get(apiUrl)

	if err != nil {
		fmt.Println("Error calling API:", err.Error())
		return
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err.Error())
		return
	}

	var responseFromServer ResponseFromServer
	json.Unmarshal(body, &responseFromServer)

	// loop over employees
	for _, employee := range responseFromServer.Data {
		fmt.Printf("Employee: %+v\n", employee)
	}
}
