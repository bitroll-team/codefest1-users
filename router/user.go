package router

import (
	"bitroll/codefest1-users/config"
	"bitroll/codefest1-users/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) Register(ctx *gin.Context) {
	// validate

	var req model.ReqRegister
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.MsgBadRequest())
		return
	}

	if err := r.validator.Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, model.MsgValidationErr(err.Error()))
		return
	}

	if err := r.ctrl.Register(req); err != nil {
		log.Println(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.MsgIntServerErr())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "User registered"})
}

func (r *Router) RegisterTeacher(ctx *gin.Context) {

	// validate

	var req model.ReqRegisterTeacher
	if err := ctx.BindJSON(&req); err != nil {
		log.Println(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.MsgBadRequest())
		return
	}
	log.Println("wowo2")

	// validate access token

	err, _ := ValidateToken(req.Token, []byte(config.Cfg.Secret))
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid access token"})
		return
	}

	// validate struct

	if err := r.validator.Struct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, model.MsgValidationErr(err.Error()))
		return
	}

	if err := r.ctrl.RegisterTeacher(req); err != nil {
		log.Println(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.MsgIntServerErr())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "User registered"})
}
