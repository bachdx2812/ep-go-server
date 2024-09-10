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
		{
			ID:             2,
			EmployeeName:   "Maria Onizuka",
			EmployeeSalary: 2000,
			EmployeeAge:    25,
			ProfileImage:   "https://dummyimage.com/600x400/000/fff.png&text=Maria+Onizuka",
		},
		{
			ID:             3,
			EmployeeName:   "Emily",
			EmployeeSalary: 3000,
			EmployeeAge:    28,
			ProfileImage:   "https://dummyimage.com/600x400/000/fff.png&text=Emily",
		},
		{
			ID:             4,
			EmployeeName:   "Tom",
			EmployeeSalary: 4000,
			EmployeeAge:    35,
			ProfileImage:   "https://dummyimage.com/600x400/000/fff.png&text=Tom",
		},
		{
			ID:             5,
			EmployeeName:   "John Snow",
			EmployeeSalary: 5000,
			EmployeeAge:    40,
			ProfileImage:   "https://dummyimage.com/600x400/000/fff.png&text=John+Snow",
		},
		{
			ID:             6,
			EmployeeName:   "Sarah Connor",
			EmployeeSalary: 6000,
			EmployeeAge:    32,
			ProfileImage:   "https://dummyimage.com/600x400/000/fff.png&text=Sarah+Connor",
		},
		{
			ID:             7,
			EmployeeName:   "Michael Smith",
			EmployeeSalary: 7000,
			EmployeeAge:    45,
			ProfileImage:   "https://dummyimage.com/600x400/000/fff.png&text=Michael+Smith",
		},
		{
			ID:             8,
			EmployeeName:   "Jessica Jones",
			EmployeeSalary: 8000,
			EmployeeAge:    29,
			ProfileImage:   "https://dummyimage.com/600x400/000/fff.png&text=Jessica+Jones",
		},
		{
			ID:             9,
			EmployeeName:   "Clark Kent",
			EmployeeSalary: 9000,
			EmployeeAge:    33,
			ProfileImage:   "https://dummyimage.com/600x400/000/fff.png&text=Clark+Kent",
		},
		{
			ID:             10,
			EmployeeName:   "Bruce Wayne",
			EmployeeSalary: 10000,
			EmployeeAge:    41,
			ProfileImage:   "https://dummyimage.com/600x400/000/fff.png&text=Bruce+Wayne",
		},
		{
			ID:             11,
			EmployeeName:   "Diana Prince",
			EmployeeSalary: 11000,
			EmployeeAge:    28,
			ProfileImage:   "https://dummyimage.com/600x400/000/fff.png&text=Diana+Prince",
		},
		{
			ID:             12,
			EmployeeName:   "Peter Parker",
			EmployeeSalary: 12000,
			EmployeeAge:    22,
			ProfileImage:   "https://dummyimage.com/600x400/000/fff.png&text=Peter+Parker",
		},
		{
			ID:             13,
			EmployeeName:   "Tony Stark",
			EmployeeSalary: 13000,
			EmployeeAge:    45,
			ProfileImage:   "https://dummyimage.com/600x400/000/fff.png&text=Tony+Stark",
		},
		{
			ID:             14,
			EmployeeName:   "Natasha Romanoff",
			EmployeeSalary: 14000,
			EmployeeAge:    35,
			ProfileImage:   "https://dummyimage.com/600x400/000/fff.png&text=Natasha+Romanoff",
		},
		{
			ID:             15,
			EmployeeName:   "Steve Rogers",
			EmployeeSalary: 15000,
			EmployeeAge:    38,
			ProfileImage:   "https://dummyimage.com/600x400/000/fff.png&text=Steve+Rogers",
		},
		{
			ID:             16,
			EmployeeName:   "Wade Wilson",
			EmployeeSalary: 16000,
			EmployeeAge:    34,
			ProfileImage:   "https://dummyimage.com/600x400/000/fff.png&text=Wade+Wilson",
		},
		{
			ID:             17,
			EmployeeName:   "Barry Allen",
			EmployeeSalary: 17000,
			EmployeeAge:    27,
			ProfileImage:   "https://dummyimage.com/600x400/000/fff.png&text=Barry+Allen",
		},
		{
			ID:             18,
			EmployeeName:   "Hal Jordan",
			EmployeeSalary: 18000,
			EmployeeAge:    39,
			ProfileImage:   "https://dummyimage.com/600x400/000/fff.png&text=Hal+Jordan",
		},
		{
			ID:             19,
			EmployeeName:   "Arthur Curry",
			EmployeeSalary: 19000,
			EmployeeAge:    33,
			ProfileImage:   "https://dummyimage.com/600x400/000/fff.png&text=Arthur+Curry",
		},
		{
			ID:             20,
			EmployeeName:   "T'Challa",
			EmployeeSalary: 20000,
			EmployeeAge:    37,
			ProfileImage:   "https://dummyimage.com/600x400/000/fff.png&text=T'Challa",
		},
		{
			ID:             21,
			EmployeeName:   "Bucky Barnes",
			EmployeeSalary: 21000,
			EmployeeAge:    32,
			ProfileImage:   "https://dummyimage.com/600x400/000/fff.png&text=Bucky+Barnes",
		},
		{
			ID:             22,
			EmployeeName:   "Sam Wilson",
			EmployeeSalary: 22000,
			EmployeeAge:    35,
			ProfileImage:   "https://dummyimage.com/600x400/000/fff.png&text=Sam+Wilson",
		},
		{
			ID:             23,
			EmployeeName:   "Wanda Maximoff",
			EmployeeSalary: 23000,
			EmployeeAge:    29,
			ProfileImage:   "https://dummyimage.com/600x400/000/fff.png&text=Wanda+Maximoff",
		},
		{
			ID:             24,
			EmployeeName:   "Vision",
			EmployeeSalary: 24000,
			EmployeeAge:    36,
			ProfileImage:   "https://dummyimage.com/600x400/000/fff.png&text=Vision",
		},
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
