package controller

import (
	"bitroll/codefest1-users/model"
	"context"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (ctrl *Controller) Register(req model.ReqRegister) error {

	// pass hash
	// TODO: Use secret

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), 8)
	if err != nil {
		return err
	}

	// write

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	query := `
		INSERT INTO users (role, username, email, full_name, password_hash)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err = ctrl.DB.ExecContext(
		ctx,
		query,
		"STUDENT",
		req.Username,
		req.Email,
		req.Fullname,
		string(hashed),
	)
	if err != nil {
		return err
	}

	return nil
}

func (ctrl *Controller) RegisterTeacher(req model.ReqRegisterTeacher) error {

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), 8)
	if err != nil {
		return err
	}

	// write

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	query := `
		INSERT INTO users (role, username, email, full_name, password_hash)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err = ctrl.DB.ExecContext(
		ctx,
		query,
		"TEACHER",
		req.Username,
		req.Email,
		req.Fullname,
		string(hashed),
	)
	if err != nil {
		return err
	}

	return nil
}

func (ctrl *Controller) SearchUser(req model.ReqSearchUser) ([]model.UserSearchResult, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
	SELECT uuid, username FROM Users WHERE username LIKE CONCAT('%', CAST($1 as VARCHAR(50)), '%')
	`

	rows, err := ctrl.DB.QueryContext(
		ctx,
		query,
		req.Query,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// iterate results

	var users []model.UserSearchResult

	for rows.Next() {
		var user model.UserSearchResult
		if err := rows.Scan(&user.UUID, &user.Username); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
