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
		INSERT INTO users (username, email, full_name, password_hash)
		VALUES ($1, $2, $3, $4)
	`

	_, err = ctrl.DB.ExecContext(
		ctx,
		query,
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
