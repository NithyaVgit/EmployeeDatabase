package respoistory

import (
	"errors"
)

/*
Employee Insert Function
Arguments : Value is Interface
Return : Insert Error
*/
func InsertEmployeeDetails(value *Employee) error {
	return MysqlDB.Create(&value).Error
}

/*Employee Find By Email Function
Arguments : Employee Email in String
Return : Employee Details and Find Error
*/

func FindEmployeeByEmail(email string) (*Employee, error) {
	var employee Employee
	tx := MysqlDB.Model(&Employee{}).Where("email=?", email).Find(&employee)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, errors.New("No Employee Record")
	}
	return &employee, nil
}

/*Employee FindAll Function
Return : All Employee Details and Find Error
*/

func FindAllEmployee() ([]Employee, error) {
	var employee []Employee
	tx := MysqlDB.Model(&Employee{}).Find(&employee)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, errors.New("No Employee Record")
	}
	return employee, nil
}

/*
Employee Update Function
Arguments : Value is Interface && key String
Return : Update Error
*/

func UpdateEmployeeDetailsBasedOnEmail(value interface{}, email string) error {
	return MysqlDB.Model(&Employee{}).Where("email=?", email).Updates(value).Error
}
