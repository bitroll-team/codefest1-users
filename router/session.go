package router

import (
	"bitroll/codefest1-users/config"
	"bitroll/codefest1-users/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const EXP_ACCESS_COOKIE = 60 * 4

func (r *Router) Login(ctx *gin.Context) {
	// validate

	var req model.ReqLogin
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.MsgBadRequest())
		return
	}

	if err := r.validator.Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, model.MsgValidationErr(err.Error()))
		return
	}

	userId, err := r.ctrl.Login(req)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.MsgIntServerErr())
		return
	}

	// TODO: Create token

	token, _, err := CreateToken(TokenInfo{UserID: uuid.MustParse(userId)}, []byte(config.Cfg.Secret))
	if err != nil {
		log.Println(err)
		return
	}

	// send

	ctx.JSON(http.StatusOK, gin.H{
		"msg":   "User logged in",
		"token": token,
	})
}
