package repository

import (
	"context"
	"database/sql"
)

func (r *Repository) GetTestById(ctx context.Context, input GetTestByIdInput) (output GetTestByIdOutput, err error) {
	err = r.Db.QueryRowContext(ctx, "SELECT name FROM test WHERE id = $1", input.Id).Scan(&output.Name)
	if err != nil {
		return
	}
	return
}

func (r *Repository) GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (*GetUserByPhoneNumberOutput, error) {
	var output GetUserByPhoneNumberOutput

	err := r.Db.QueryRowContext(ctx, "SELECT id, full_name, phone_number FROM users WHERE phone_number = $1", phoneNumber).
		Scan(&output.Id, &output.FullName, &output.PhoneNumber)

	if err == sql.ErrNoRows {
		// Handle the case when the user is not found
		return nil, nil
	} else if err != nil {
		// Handle other errors
		return nil, err
	}

	return &output, nil
}

func (r *Repository) GetUserById(ctx context.Context, id int) (*User, error) {
	var output User

	err := r.Db.QueryRowContext(ctx, "SELECT full_name, phone_number FROM users WHERE id = $1", id).
		Scan(&output.FullName, &output.PhoneNumber)

	if err == sql.ErrNoRows {
		// Handle the case when the user is not found
		return nil, nil
	} else if err != nil {
		// Handle other errors
		return nil, err
	}

	return &output, nil
}

func (r *Repository) CreateUserInput(ctx context.Context, input *CreateUserInput) (int, error) {
	var id int

	err := r.Db.QueryRowContext(ctx, "INSERT INTO users (full_name, password, phone_number) VALUES ($1, $2, $3) RETURNING id",
		input.FullName, input.Password, input.PhoneNumber).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *Repository) UpdateUserById(ctx context.Context, input *UpdateUserByIdInput, id int) (*UpdateUserByIdOutput, error) {
	var output UpdateUserByIdOutput

	err := r.Db.QueryRowContext(ctx, "UPDATE users SET full_name = $1, phone_number = $2 WHERE id = $3 RETURNING id, full_name, phone_number",
		input.FullName, input.PhoneNumber, id).Scan(&output.Id, &output.FullName, &output.PhoneNumber)
	if err != nil {
		return nil, err
	}
	return &output, nil
}
