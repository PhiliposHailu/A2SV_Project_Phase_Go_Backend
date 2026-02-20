package main

import (
	"github.com/gin-gonic/gin"
	"github.com/philipos/api/controllers"
	"github.com/philipos/api/router"
)

func main() {
	controllers.LoadData()
	r := gin.Default()

	router.TaskRouters(r)

	r.Run("localhost:3000")

}
