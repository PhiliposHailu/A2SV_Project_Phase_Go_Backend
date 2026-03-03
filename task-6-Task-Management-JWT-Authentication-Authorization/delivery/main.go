package main

import (
	"github.com/gin-gonic/gin"
	"github.com/philipos/api/usecase"
	"github.com/philipos/api/delivery/router"
)

func main() {
	data.ConnectDB()
	r := gin.Default()

	router.TaskRouters(r)

	r.Run("localhost:3000")
}
