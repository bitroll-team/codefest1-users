package controller

import (
	"bitroll/codefest1-users/model"
	"context"
	"golang.org/x/crypto/bcrypt"
	"time"
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
