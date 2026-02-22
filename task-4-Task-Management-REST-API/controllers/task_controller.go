package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/philipos/api/data"
	"github.com/philipos/api/models"
)

// CRUD operations
func GetAllTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message: ": data.GetAllTasksService()})
}

func GetTask(c *gin.Context) {
	id := c.Param("id")

	task, err := data.GetTaskService(id)
	if err != nil {
		c.IndentedJSON(http.StatusOK, gin.H{"message": task})
		return
	}
	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
}

func CreateTask(c *gin.Context) {
	var newTask models.Task
	if err := c.BindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message: ": "failed to decerialize to obj"})
		return
	}
	data.CreateTaskService(newTask)
}

func UpdateTask(c *gin.Context) {
	var newTask models.Task
	if err := c.BindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message: ": "failed to decerialize to obj"})
		return
	}
	id := c.Param("id")
	err := data.UpdateTaskService(id, newTask)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "updated successfuly."})

}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	err := data.DeleteTaskService(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "deleted successfuly."})
}
