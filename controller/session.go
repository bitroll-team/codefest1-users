package controller

import (
	"bitroll/codefest1-users/model"
	"context"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (ctrl *Controller) Login(req model.ReqLogin) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	query := "SELECT id,password_hash FROM users WHERE email = $1"

	row := ctrl.DB.QueryRowContext(
		ctx,
		query,
		req.Email,
	)

	var userId string
	var password_hash string

	if err := row.Scan(&userId, &password_hash); err != nil {
		return "", err
	}

	// compare

	if err := bcrypt.CompareHashAndPassword([]byte(password_hash), []byte(req.Password)); err != nil {
		return "", err
	}

	return userId, nil
}
