package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/philipos/api/models"
)

var tasks = []models.Task{}

const dbFile = "data/task_service.json"

func saveData() {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling data:", err)
		return
	}

	err = os.WriteFile(dbFile, data, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
	}
}

func LoadData() {
	data, err := os.ReadFile(dbFile)
	if err != nil {
		return
	}
	fmt.Println("before", tasks)

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		fmt.Println("Error unmarshalling data:", err)
	}
	fmt.Println("after", tasks)
}

// CRUD operations
func GetAllTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message: ": tasks})
}

func GetTask(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message: ": "id should be an int value!"})
		return
	}
	if len(tasks) >= intId {
		c.IndentedJSON(http.StatusOK, gin.H{"message: ": tasks[intId-1]})
		return
	}

	c.IndentedJSON(http.StatusBadRequest, gin.H{"message: ": "no task with given id."})

}

func CreateTask(c *gin.Context) {
	var newTask models.Task
	if err := c.BindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message: ": "failed to decerialize to obj"})
		return
	}

	tasks = append(tasks, newTask)
	saveData()

}

func UpdateTask(c *gin.Context) {
	var newTask models.Task
	if err := c.BindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message: ": "failed to decerialize to obj"})
		return
	}

	id := c.Param("id")
	intId, err := strconv.Atoi(id)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message: ": "id should be an int value!"})
		return
	}

	if len(tasks) >= intId {
		tasks[intId-1] = newTask
		c.IndentedJSON(http.StatusOK, gin.H{"message: ": "Task updated"})
		saveData()
		return
	}

}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message: ": "id should be an int value!"})
		return
	}

	if len(tasks) >= intId {
		tasks = append(tasks[:intId-1], tasks[intId:]...)
		c.IndentedJSON(http.StatusOK, gin.H{"message: ": "Task Deleted."})
		saveData()
		return
	}
}
