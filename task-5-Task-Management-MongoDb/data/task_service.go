package data

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/philipos/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var TasksCollection *mongo.Collection

func ConnectDB() {
	uri := "mongodb://localhost:27017"

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	// Check connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Connection failed! Is MongoDB running? Error:", err)
	}

	fmt.Println("✅ Successfully connected to MongoDB!")

	TasksCollection = client.Database("taskdb").Collection("tasks")
}

func GetAllTasksService() ([]models.Task, error) {
	foundTasks := []models.Task{}

	filter := bson.M{}
	cursor, err := TasksCollection.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}

	// DECODE bson in to a struct 
	err = cursor.All(context.TODO(), &foundTasks)
	if err != nil {
		return nil, err
	}

	return foundTasks, nil
}

func GetTaskService(id primitive.ObjectID) (*models.Task, error) {
	var foundTask models.Task
	filter := bson.M{"_id": id}

	err := TasksCollection.FindOne(context.TODO(), filter).Decode(&foundTask)
	if err != nil {
		return nil, err
	}
	return &foundTask, err
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

func DeleteTaskService(id string) error {
	intId, err := strconv.Atoi(id)

	if err != nil {
		return fmt.Errorf("could not convert id %s to int value.", id)
	}

	for i, task := range tasks {
		if task.ID == intId {
			tasks = append(tasks[:i], tasks[i+1:]...)
			saveData()
			return nil
		}
	}
	return fmt.Errorf("task with id: %s does not exist", id)
}
