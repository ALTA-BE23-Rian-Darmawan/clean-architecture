package users

import "time"

type User struct {
	UserID      uint
	FullName    string
	Email       string
	Password    string
	PhoneNumber string
	Address     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type DataUserInterface interface {
	CreateAccount(account User) error
	AccountByEmail(email string) (*User, error)
}

type ServiceUserInterface interface {
	RegistrasiAccount(accounts User) error
	LoginAccount(email string, password string) (data *User, token string, err error)
}
