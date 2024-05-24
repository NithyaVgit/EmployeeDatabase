package handler

import (
	"encoding/json"
	"go/employee/attendance/domain"
	"go/employee/attendance/output"
	"go/employee/attendance/service"
	"net/http"
)

func EmployeeInsertHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody domain.EmployeeClientRequest
	jsonData := json.NewDecoder(r.Body)
	bodyErr := jsonData.Decode(&requestBody)
	if bodyErr != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		err := output.ErrorRequest(
			"Employee Request Body Value is Empty",
			"Bad Request",
			http.StatusBadRequest,
		)
		json.NewEncoder(w).Encode(err)
		return
	}
	defer r.Body.Close()

	resp, err := service.EmployeeInsertService(requestBody)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(err.StatusCode)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	json.NewEncoder(w).Encode(resp)
}

func GetEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("employee_email")

	resp, err := service.GetEmployeeService(query)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(err.StatusCode)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	json.NewEncoder(w).Encode(resp)
}

func GetAllEmployeeHandler(w http.ResponseWriter, r *http.Request) {

	resp, err := service.GetAllEmployeeService()

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(err.StatusCode)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	json.NewEncoder(w).Encode(resp)
}

func EmployeeUpdateHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody domain.EmployeeClientRequest
	jsonData := json.NewDecoder(r.Body)
	bodyErr := jsonData.Decode(&requestBody)
	if bodyErr != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		err := output.ErrorRequest(
			"Employee Request Body Value is Empty",
			"Bad Request",
			http.StatusBadRequest,
		)
		json.NewEncoder(w).Encode(err)
		return
	}
	defer r.Body.Close()

	resp, err := service.UpdateEmployeeService(requestBody)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(err.StatusCode)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	json.NewEncoder(w).Encode(resp)
}

func EmployeecsvInsertHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10 MB max
	if err != nil {
		reserr := output.ErrorRequest("File Size Exceed", "Bad Request", http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(reserr.StatusCode)
		json.NewEncoder(w).Encode(reserr)
		return
	}

	// Retrieve the file from the form data
	file, fileHandler, err := r.FormFile("employeecsv")

	if err != nil {
		reserr := output.ErrorRequest("Unable to retrieve file from form", "Bad Request", http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(reserr.StatusCode)
		json.NewEncoder(w).Encode(reserr)
		return
	}
	defer file.Close()

	resp, reserr := service.EmployeeInsertCSVService(file, fileHandler.Filename)

	if reserr != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(reserr.StatusCode)
		json.NewEncoder(w).Encode(reserr)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	json.NewEncoder(w).Encode(resp)

}

func EmployeeCsvDownloadHandler(w http.ResponseWriter, r *http.Request) {

	resp, err := service.EmployeeDownloadtCSVService()

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(err.StatusCode)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	json.NewEncoder(w).Encode(resp)
}
