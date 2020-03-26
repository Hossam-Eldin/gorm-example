package app

import (
	controllers "gorm-example/src/controllers"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func Routes() {
	router.GET("/", controllers.Test)
	router.Run(":3000")
}
