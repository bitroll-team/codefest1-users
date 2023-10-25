package controller

import (
	"bitroll/codefest1-users/model"
	"context"
	"time"

	"github.com/google/uuid"
)

func (ctrl *Controller) FollowUser(req model.ReqFollowUser, userId uuid.UUID) error {

	// write

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	query := `
		INSERT INTO followers (follower_user_uuid, followed_user_uuid)
		VALUES ($1, $2)
	`
	_, err := ctrl.DB.ExecContext(
		ctx,
		query,
		userId,
		req.OtherUserId,
	)
	if err != nil {
		return err
	}

	return nil
}

func (ctrl *Controller) UnFollowUser(req model.ReqFollowUser, userId uuid.UUID) error {

	// write

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	query := `
		DELETE FROM followers WHERE follower_user_uuid = $1 AND followed_user_uuid = $2
	`
	_, err := ctrl.DB.ExecContext(
		ctx,
		query,
		userId,
		req.OtherUserId,
	)
	if err != nil {
		return err
	}

	return nil
}
