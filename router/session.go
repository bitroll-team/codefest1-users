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

	userId, username, err := r.ctrl.Login(req)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.MsgIntServerErr())
		return
	}

	// Create token

	token, _, err := CreateToken(TokenInfo{UserID: uuid.MustParse(userId)}, []byte(config.Cfg.Secret))
	if err != nil {
		log.Println(err)
		return
	}

	// send

	ctx.JSON(http.StatusOK, gin.H{
		"msg":      "User logged in",
		"username": username,
		"token":    token,
	})
}

func (r *Router) Challenge(ctx *gin.Context) {

	// validate

	var req model.ReqChallenge
	if err := ctx.BindJSON(&req); err != nil {
		log.Println(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.MsgBadRequest())
		return
	}

	// validate access token

	err, _ := ValidateToken(req.Token, []byte(config.Cfg.Secret))
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid access token"})
		return
	}

	// send

	ctx.JSON(http.StatusOK, gin.H{"msg": "Auth success"})
}
