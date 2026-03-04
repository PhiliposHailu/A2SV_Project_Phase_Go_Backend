package repository

import (
	"context"
	"fmt"

	"github.com/philipos/api/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type taskRepository struct {
	database   *mongo.Database
	collection string
}

func NewTaskRepository(db *mongo.Database, collection string) domain.TaskRepository {
	return &taskRepository{
		database:   db,
		collection: collection,
	}
}

func (r *taskRepository) FetchAll() ([]domain.Task, error) {
	var tasks []domain.Task
	collection := r.database.Collection(r.collection)

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	err = cursor.All(context.TODO(), &tasks)
	return tasks, err
}

func (r *taskRepository) GetByID(id string) (*domain.Task, error) {
	var task domain.Task
	collection := r.database.Collection(r.collection)

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *taskRepository) Create(task *domain.Task) error {
	collection := r.database.Collection(r.collection)

	_, err := collection.InsertOne(context.TODO(), task)
	return err
}

func (r *taskRepository) Update(id string, task *domain.Task) error {
	collection := r.database.Collection(r.collection)
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	updateData := bson.M{}

	updateData["due_date"] = task.DueDate
	if task.Title != "" {
		updateData["title"] = task.Title
	}
	if task.Description != "" {
		updateData["description"] = task.Description
	}
	if task.Status != "" {
		updateData["status"] = task.Status
	}

	if len(updateData) == 0 {
		return fmt.Errorf("no fields provided")
	}

	update := bson.M{
		"$set": updateData,
	}

	_, err = collection.UpdateOne(context.TODO(), bson.M{"_id": objID}, update)
	return err
}

func (r *taskRepository) Delete(id string) error {
	collection := r.database.Collection(r.collection)
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	return err
}
