package router

import (
	"bitroll/codefest1-users/config"
	"bitroll/codefest1-users/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) FollowUser(ctx *gin.Context) {
	// validate

	var req model.ReqFollowUser
	if err := ctx.BindJSON(&req); err != nil {
		log.Println(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.MsgBadRequest())
		return
	}
	if err := r.validator.Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, model.MsgValidationErr(err.Error()))
		return
	}

	// validate access token

	err, claims := ValidateToken(req.Token, []byte(config.Cfg.Secret))
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized,
			gin.H{"message": "Invalid access token"})
		return
	}
	userId := claims.UserID

	if err := r.ctrl.FollowUser(req, userId); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.MsgIntServerErr())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "Following user"})
}
