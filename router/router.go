package router

import (
	"bitroll/codefest1-users/controller"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Router struct {
	*gin.Engine
	ctrl      *controller.Controller
	validator *validator.Validate
}

// TODO: Move config

func SetupRouter(ctrl *controller.Controller) Router {

	var r Router
	r.Engine = gin.Default()
	r.ctrl = ctrl
	r.validator = validator.New()

	// routes

	base := r.Group("/api/v1")

	user := base.Group("/user")
	user.POST("/register", r.Register)

	sess := base.Group("/session")
	sess.POST("/login", r.Login)
	//user.POST("/refresh", r.Refresh)

	return r
}
