package models

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/ikilonchic/WEB_LAB-3-4/pkg/validation"
)

// User ...
type User struct {
	ID              string     `json:"id"`
	UserName		string	   `json:"username"`
	Email    		string     `json:"email"`
	Password 		string     `json:"password"`
	Number			string	   `json:"number"`
	Male 	 		bool	   `json:"male"`
	Country			string	   `json:"country"`
	DateOfBirth	    time.Time  `json:"dob"`
	About			string	   `json:"about"`
}

// Validate ...
func (u *User) Validate() bool {
	return validation.IsEmail(u.Email) && validation.IsPassword(u.Password)
}

// BeforeCreate ...
func (u *User) BeforeCreate(*gorm.DB) error {
	if len(u.Password) > 0 {
		encod, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
		if err != nil {
			return err
		}

		u.Password = string(encod)
	}

	u.ID = uuid.New().String()

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