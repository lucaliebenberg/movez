package types

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const(
	bcryptCost = 12
	minFirstNameLen = 2
	minLastNameLen = 2
	minPasswordLen = 7
)

type Location struct {
	Lat string
	Long string
}

type UpdateUserParams struct {
	FirstName string 
	LastName string
}

type CreateUserParams struct {
	FirstName string 
	LastName string
	Password string
	Email string
	LocationHistory []Location
}

type User struct {
	ID int
	FirstName string 
	LastName string
	EncryptedPassword string
	Email string
	LocationHistory []Location
}

// Functions
// func (params CreateUserParams) bool {
// 	fmt.Println("Validating newly created user")
// 	return false
// }

func isEmailValid(email string) bool {
	fmt.Println("Validating email address")
	// emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return false
}

func isValidPassword(encpwd, pw string) bool {
	fmt.Println("Validating newly created user")
	return bcrypt.CompareHashAndPassword([]byte(encpwd), []byte(pw)) == nil
}

// func NewUserFromParams(params CreateUserParams) (*User, error) {
// 	fmt.Println("New user created from params")
// 	encpw, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcryptCost)
// 	return &User{
// 		FirstName: params.FirstName,
// 		LastName: params.LastName,
// 		Email: params.Email,
// 		EncryptedPassword: string(encpw),
// 	}
// }