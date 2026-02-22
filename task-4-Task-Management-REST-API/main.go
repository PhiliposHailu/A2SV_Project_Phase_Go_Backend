package main

import (
	"github.com/gin-gonic/gin"
	"github.com/philipos/api/data"
	"github.com/philipos/api/router"
)

func main() {
	data.LoadData()
	r := gin.Default()

	router.TaskRouters(r)

	r.Run("localhost:3000")

}
