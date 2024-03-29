// This file contains the interfaces for the repository layer.
// The repository layer is responsible for interacting with the database.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package repository

import "context"

type RepositoryInterface interface {
	GetTestById(ctx context.Context, input GetTestByIdInput) (output GetTestByIdOutput, err error)
	GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (output *GetUserByPhoneNumberOutput, err error)
	GetUserById(ctx context.Context, id int) (output *User, err error)

	CreateUserInput(ctx context.Context, input *CreateUserInput) (int, error)
	UpdateUserById(ctx context.Context, input *UpdateUserByIdInput, id int) (output *UpdateUserByIdOutput, err error)
}
