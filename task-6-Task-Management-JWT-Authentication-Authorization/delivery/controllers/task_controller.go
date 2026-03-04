package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/philipos/api/domain" 
)

type TaskController struct {
	taskUsecase domain.TaskUsecase
}

func NewTaskController(us domain.TaskUsecase) *TaskController {
	return &TaskController{
		taskUsecase: us,
	}
}


func (h *TaskController) FetchAll(c *gin.Context) {
	tasks, err := h.taskUsecase.FetchAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (h *TaskController) GetByID(c *gin.Context) {
	id := c.Param("id")
	task, err := h.taskUsecase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (h *TaskController) Create(c *gin.Context) {
	var task domain.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.taskUsecase.Create(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Task created successfully"})
}

func (h *TaskController) Update(c *gin.Context) {
	id := c.Param("id")
	var task domain.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.taskUsecase.Update(id, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}

func (h *TaskController) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.taskUsecase.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}