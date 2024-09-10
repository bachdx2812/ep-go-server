package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

const (
	apiUrl     = "https://dummy.restapiexample.com/api/v1/employees"
	maxWorkers = 5 // Limit the number of concurrent workers
)

var dummyResponseFromServer = ResponseFromServer{
	Status: "success",
	Data: []Employee{
		{
			ID:             1,
			EmployeeName:   "John Snow",
			EmployeeSalary: 1000,
			EmployeeAge:    30,
			ProfileImage:   "https://dummyimage.com/600x400/000/fff.png&text=John+Snow",
		},
		// ... (other employees)
		{
			ID:             25,
			EmployeeName:   "Scott Lang",
			EmployeeSalary: 25000,
			EmployeeAge:    40,
			ProfileImage:   "https://dummyimage.com/600x400/000/fff.png&text=Scott+Lang",
		},
	},
}

// Employee struct to hold individual employee data
type Employee struct {
	ID             int    `json:"id"`
	EmployeeName   string `json:"employee_name"`
	EmployeeSalary int    `json:"employee_salary"`
	EmployeeAge    int    `json:"employee_age"`
	ProfileImage   string `json:"profile_image"`
}

// ResponseFromServer represents the structure of the API response
type ResponseFromServer struct {
	Status string     `json:"status"`
	Data   []Employee `json:"data"`
}

func main() {
	responseFromServer, err := fetchEmployees()
	if err != nil {
		fmt.Println("Error fetching employees:", err.Error())
		return
	}

	fmt.Println("START ASSIGNING JOBS")
	assignJobsWithWorkerPool(responseFromServer.Data)
}

func fetchEmployees() (ResponseFromServer, error) {
	var responseFromServer ResponseFromServer

	response, err := http.Get(apiUrl)
	if err != nil {
		fmt.Println("Error calling API, using dummy data:", err.Error())
		return dummyResponseFromServer, nil
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body, using dummy data:", err.Error())
		return dummyResponseFromServer, nil
	}

	// Check if the response status is OK
	if response.StatusCode != http.StatusOK {
		fmt.Println("API response not OK, using dummy data")
		return dummyResponseFromServer, nil
	}

	if err := json.Unmarshal(body, &responseFromServer); err != nil {
		fmt.Println("Error unmarshaling response, using dummy data:", err.Error())
		return dummyResponseFromServer, nil
	}

	fmt.Printf("Fetched %d employees from API\n", len(responseFromServer.Data))
	return responseFromServer, nil
}

func worker(id int, jobs <-chan Employee, wg *sync.WaitGroup) {
	for employee := range jobs {
		time.Sleep(5 * time.Second)
		averageSalary := employee.EmployeeSalary / employee.EmployeeAge
		fmt.Printf("Worker %d processed employee %s, average salary per age: %d\n", id, employee.EmployeeName, averageSalary)
		wg.Done()
	}
}

func assignJobsWithWorkerPool(employees []Employee) {
	var wg sync.WaitGroup

	jobs := make(chan Employee, len(employees))

	for i := 1; i <= maxWorkers; i++ {
		go worker(i, jobs, &wg)
	}

	for _, employee := range employees {
		wg.Add(1)
		jobs <- employee
	}

	close(jobs)

	wg.Wait()
	fmt.Println("All jobs completed")
}
