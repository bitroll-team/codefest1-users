package controller

import (
	"bitroll/codefest1-users/model"
	"context"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (ctrl *Controller) Login(req model.ReqLogin) (*model.ResLogin, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	query := "SELECT password_hash,uuid,username,role FROM users WHERE email = $1"

	row := ctrl.DB.QueryRowContext(
		ctx,
		query,
		req.Email,
	)

	var res model.ResLogin
	var passwordHash string

	if err := row.Scan(&passwordHash, &res.UserId, &res.Username, &res.Role); err != nil {
		return nil, err
	}

	// compare

	if err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(req.Password)); err != nil {
		return nil, err
	}

	return &res, nil
}
