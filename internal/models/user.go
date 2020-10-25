package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/ikilonchic/WEB_LAB-3-4/pkg/validation"
)

// User ...
type User struct {
	ID              int        `json:"id"`
	UserName		string	   `json:"username"`
	Email    		string     `json:"email"`
	Password 		string     `json:"password"`
	Number			string	   `json:"number"`
	Male 	 		bool	   `json:"male"`
	DateOfBirth	    time.Time  `json:"dof"`
	About			string	   `json:"about"`
}

// Validate ...
func (u *User) Validate() bool {
	return validation.IsEmail(u.Email) && validation.IsPassword(u.Password)
}

// BeforeCreate ...
func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		encod, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
		if err != nil {
			return err
		}

		u.Password = string(encod)
	}

	return nil
}

// Sanitize ...
func (u *User) Sanitize() {
	u.Password = ""
}

// ComparePassword ...
func (u *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}