package data

import (
	"encoding/json"
	"fmt"
	"strconv"
	"os"
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

func GetAllTasksService() []models.Task {
	return tasks
}

func GetTaskService(id string) (*models.Task, error) {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("id not correct id: %s", id)
	}

	for _, task := range tasks {
		if task.ID == intId {
			return &task, nil
		}
	}
	return nil, fmt.Errorf("task with id: %s does not exist", id)
}

func CreateTaskService(newTask models.Task) {
	tasks = append(tasks, newTask)
	saveData()
}

func UpdateTaskService(id string, updatedTask models.Task) error {
	intId, err := strconv.Atoi(id)

	if err != nil {
		return fmt.Errorf("id not correct id: %s", id)
	}

	for i, task := range tasks {
		if task.ID == intId {
			tasks[i] = updatedTask
			saveData()
			return nil
		}
	}
	return fmt.Errorf("task with id: %s does not exist", id)
}

func DeleteTaskService(id string) (error){
	intId, err := strconv.Atoi(id)

	if err != nil {
		return fmt.Errorf("could not convert id %s to int value.", id)
	}

	for i, task := range tasks {
		if task.ID == intId {
			tasks = append(tasks[:i], tasks[i + 1:]...)
			saveData()
			return nil
		}
	}
	return fmt.Errorf("task with id: %s does not exist", id)
}
