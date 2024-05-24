package respoistory

import "time"

type Employee struct {
	ID          uint      `gorm:"column:id;pirmary_key;type:int(25);autoIncrement"`
	Email       string    `gorm:"column:email;unique;type:varchar(45);not null"`
	FirstName   string    `gorm:"column:first_name;type:varchar(30)"`
	LastName    string    `gorm:"column:last_name;type:varchar(30)"`
	PhoneNumber string    `gorm:"column:phone_number;unique;type:varchar(10)"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}
