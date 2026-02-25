package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/philipos/api/data"
	"github.com/philipos/api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CRUD operations
func GetAllTasks(c *gin.Context) {
	foundTasks, err := data.GetAllTasksService()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": foundTasks})
}

func GetTask(c *gin.Context) {
	id := c.Param("id")
	// parse it in to a primitive object type 
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	task, err := data.GetTaskService(objId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": task})
}

func CreateTask(c *gin.Context) {
	var newTask models.Task
	if err := c.BindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error with given task"})
		return
	}
	data.CreateTaskService(newTask)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Task created successfuly."})
}

func UpdateTask(c *gin.Context) {
	var newTask models.Task
	if err := c.BindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "failed to decerialize to obj"})
		return
	}
	id := c.Param("id")
	err := data.UpdateTaskService(id, newTask)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "updated successfuly."})

}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	err := data.DeleteTaskService(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "deleted successfuly."})
}
