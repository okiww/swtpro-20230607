// This file contains types that are used in the repository layer.
package repository

type GetTestByIdInput struct {
	Id string
}

type GetTestByIdOutput struct {
	Name string
}

type CreateUserInput struct {
	FullName    string
	Password    string
	PhoneNumber string
}

type GetUserByPhoneNumberOutput struct {
	Id          int
	FullName    string
	PhoneNumber string
}

type UpdateUserByIdInput struct {
	FullName    string
	PhoneNumber string
}

type UpdateUserByIdOutput struct {
	Id          int
	FullName    string
	PhoneNumber string
}
