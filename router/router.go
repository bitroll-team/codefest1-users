package router

import (
	"bitroll/codefest1-users/controller"

	"github.com/gin-contrib/cors"
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

	// cors

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowCredentials = true
	r.Use(cors.New(config))

	// routes

	base := r.Group("/api/v1")

	user := base.Group("/user")
	user.POST("/register", r.Register)
	user.POST("/register_teacher", r.RegisterTeacher)
	user.POST("/search", r.SearchUser)
	user.POST("/follow", r.FollowUser)
	user.POST("/unfollow", r.UnFollowUser)

	sess := base.Group("/session")
	sess.POST("/login", r.Login)
	sess.POST("/challenge", r.Challenge)

	return r
}
