package service

import (
	"encoding/csv"
	"go/employee/attendance/domain"
	"go/employee/attendance/lib"
	"go/employee/attendance/output"
	mysqldb "go/employee/attendance/respoistory/Mysql"
	"go/employee/attendance/validation"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"
)

func EmployeeInsertService(req domain.EmployeeClientRequest) (*output.RestSucess, *output.Resterr) {
	valid := validation.EmployeeStructValidation(req)
	if valid != nil {
		return nil, output.ErrorRequest(
			valid.Error(),
			"Bad Request",
			http.StatusBadRequest,
		)
	}
	insertReq := mysqldb.Employee{
		Email:       req.Email,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		PhoneNumber: req.PhoneNumber,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	insertDB := mysqldb.InsertEmployeeDetails(&insertReq)
	if insertDB != nil {
		return nil, output.ErrorRequest(
			insertDB.Error(),
			"Internal Server Error",
			http.StatusInternalServerError,
		)
	}

	return output.SuccessRequest("Employee Insert Successfully", req, http.StatusOK), nil
}

func GetEmployeeService(email string) (*output.RestSucess, *output.Resterr) {
	if email == "" {
		return nil, output.ErrorRequest("Email Value is Empty", "Bad Request", http.StatusBadRequest)
	}

	dbRes, dbErr := mysqldb.FindEmployeeByEmail(email)

	if dbErr != nil {
		return nil, output.ErrorRequest(dbErr.Error(), "Internal Server Error", http.StatusInternalServerError)
	}

	return output.SuccessRequest("Get Employee Successfully", dbRes, http.StatusOK), nil
}

func GetAllEmployeeService() (*output.RestSucess, *output.Resterr) {

	dbRes, dbErr := mysqldb.FindAllEmployee()

	if dbErr != nil {
		return nil, output.ErrorRequest(dbErr.Error(), "Internal Server Error", http.StatusInternalServerError)
	}

	return output.SuccessRequest("Get All Employee Successfully", dbRes, http.StatusOK), nil
}

func UpdateEmployeeService(req domain.EmployeeClientRequest) (*output.RestSucess, *output.Resterr) {
	valid := validation.EmployeeStructValidation(req)
	if valid != nil {
		return nil, output.ErrorRequest(
			valid.Error(),
			"Bad Request",
			http.StatusBadRequest,
		)
	}

	_, dbErr := mysqldb.FindEmployeeByEmail(req.Email)

	if dbErr != nil {
		return nil, output.ErrorRequest(dbErr.Error(), "Internal Server Error", http.StatusInternalServerError)
	}

	updateReq := mysqldb.Employee{
		Email:       req.Email,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		PhoneNumber: req.PhoneNumber,
		UpdatedAt:   time.Now(),
	}

	updateDB := mysqldb.UpdateEmployeeDetailsBasedOnEmail(&updateReq, req.Email)
	if updateDB != nil {
		return nil, output.ErrorRequest(
			updateDB.Error(),
			"Internal Server Error",
			http.StatusInternalServerError,
		)
	}

	return output.SuccessRequest("Employee Update Successfully", req, http.StatusOK), nil
}

func EmployeeInsertCSVService(file multipart.File, fileName string) (*output.RestSucess, *output.Resterr) {

	if !lib.IsCSVFile(fileName) {
		return nil, output.ErrorRequest("Please Upload CSV", "Bad Request", http.StatusBadRequest)
	}

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, output.ErrorRequest("Error reading CSV:", "Internal Server Error", http.StatusInternalServerError)
	}

	// Check if the file has any records
	if len(records) == 0 {
		return nil, output.ErrorRequest("CSV File is Empty", "Bad Request", http.StatusBadRequest)
	}

	for i := 1; i < len(records); i++ {
		employee := mysqldb.Employee{
			Email:       records[i][0],
			FirstName:   records[i][1],
			LastName:    records[i][2],
			PhoneNumber: records[i][3],
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		// Check if the user already exists
		_, dbErr := mysqldb.FindEmployeeByEmail(employee.Email)
		if (dbErr == nil || dbErr.Error() == "No Employee Record") && employee.Email != "" {
			insertDB := mysqldb.InsertEmployeeDetails(&employee)
			if insertDB != nil {
				continue
			}
			continue
		}

	}
	value := strconv.Itoa(len(records)) + " Employee File Upload Success"
	return output.SuccessRequest("Upload Successfully", value, http.StatusOK), nil
}

func EmployeeDownloadtCSVService() (*output.RestSucess, *output.Resterr) {

	dbRes, dbErr := mysqldb.FindAllEmployee()

	if dbErr != nil {
		return nil, output.ErrorRequest(dbErr.Error(), "Internal Server Error", http.StatusInternalServerError)
	}

	file, err := os.Create("employee.csv")
	if err != nil {
		return nil, output.ErrorRequest("Error creating CSV file:", "Internal Server Error", http.StatusInternalServerError)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write column headers to CSV file
	columns := []string{"Email", "FirstName", "LastName", "PhoneNumber"}
	if err := writer.Write(columns); err != nil {
		return nil, output.ErrorRequest("Error writing column names to CSV:", "Internal Server Error", http.StatusInternalServerError)
	}

	// Write data rows to CSV file
	for _, row := range dbRes {
		values := []string{row.Email, row.FirstName, row.LastName, row.PhoneNumber} // Adjust value types and order as needed
		if err := writer.Write(values); err != nil {
			return nil, output.ErrorRequest("Error writing row to CSV:", "Internal Server Error", http.StatusInternalServerError)
		}
	}

	value := strconv.Itoa(len(dbRes)) + " Employee File Download Success"
	return output.SuccessRequest("Download Successfully", value, http.StatusOK), nil
}
