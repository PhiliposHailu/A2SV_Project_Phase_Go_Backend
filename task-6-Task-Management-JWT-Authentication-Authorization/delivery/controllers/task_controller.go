package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/philipos/clean-task-api/domain" // Your module path
)

// 1. The Struct
type TaskHandler struct {
	taskUsecase domain.TaskUsecase
}

// 2. The Constructor
// We pass the Gin Router AND the Usecase in from main.go
func NewTaskHandler(r *gin.Engine, us domain.TaskUsecase) {
	handler := &TaskHandler{
		taskUsecase: us,
	}

	// Define Routes right here in the delivery layer!
	r.GET("/tasks", handler.FetchAll)
	r.GET("/tasks/:id", handler.GetByID)
	r.POST("/tasks", handler.Create)
	r.PUT("/tasks/:id", handler.Update)
	r.DELETE("/tasks/:id", handler.Delete)
}

// 3. Implementing the Handlers
func (h *TaskHandler) FetchAll(c *gin.Context) {
	tasks, err := h.taskUsecase.FetchAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	task, err := h.taskUsecase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) Create(c *gin.Context) {
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

func (h *TaskHandler) Update(c *gin.Context) {
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

func (h *TaskHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.taskUsecase.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}