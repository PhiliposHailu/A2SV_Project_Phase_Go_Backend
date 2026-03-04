package usecase

import (
	"context"
	"fmt"
	"log"

	"github.com/philipos/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var TasksCollection *mongo.Collection
var UserCollection *mongo.Collection


func ConnectDB() {
	uri := "mongodb://localhost:27017"

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	// Check connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Connection failed! Is MongoDB running? Error:", err)
	}

	fmt.Println("✅ Successfully connected to MongoDB!")

	TasksCollection = client.Database("taskdb").Collection("tasks")
	UserCollection = client.Database("taskdb").Collection("users")

	indexModel := mongo.IndexModel{
        Keys:    bson.D{{Key: "username", Value: 1}},
        Options: options.Index().SetUnique(true),    
    }

    _, err = UserCollection.Indexes().CreateOne(context.TODO(), indexModel)
    if err != nil {
        log.Fatal("Could not create unique index:", err)
    }
}

func GetAllTasksService() ([]models.Task, error) {
	foundTasks := []models.Task{}

	filter := bson.M{}
	cursor, err := TasksCollection.Find(context.TODO(), filter)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.TODO())

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

func CreateTaskService(newTask models.Task) (*models.Task, error) {
	res, err := TasksCollection.InsertOne(context.TODO(), newTask)
	if err != nil {
		return nil, err
	}
	newTask.ID = res.InsertedID.(primitive.ObjectID)
	return &newTask, nil
}

func UpdateTaskService(id primitive.ObjectID, updatedTask *models.Task) (*models.Task, error) {

	filter := bson.M{
		"_id": id,
	}
	updateData := bson.M{}
	if updatedTask.Title != "" {
		updateData["title"] = updatedTask.Title
	}
	if updatedTask.Description != "" {
		updateData["description"] = updatedTask.Description
	}
	if updatedTask.Status != "" {
		updateData["status"] = updatedTask.Status
	}
	if updatedTask.DueDate != "" {
		updateData["due_date"] = updatedTask.DueDate
	}

	if len(updateData) == 0 {
		return nil, fmt.Errorf("no fields provided")
	}

	update := bson.M{
		"$set": updateData,
	}

	res, err := TasksCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	if res.MatchedCount == 0 {
		return nil, mongo.ErrNoDocuments
	}
	updatedTask.ID = id
	return updatedTask, nil
}

func DeleteTaskService(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	res, err := TasksCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}
