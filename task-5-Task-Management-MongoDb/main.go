// package main

// import (
// 	"github.com/gin-gonic/gin"
// 	"github.com/philipos/api/data"
// 	"github.com/philipos/api/router"
// )

// func main() {
// 	data.LoadData()
// 	r := gin.Default()

// 	router.TaskRouters(r)

// 	r.Run("localhost:3000")

// }

package main

import (
	"context"
	"fmt"
	"log"
	"github.com/philipos/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// DB Connection SET-UP
	uri := "mongodb://localhost:27017"

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Connection failed! Is MongoDB running? Error:", err)
	}

	fmt.Println("✅ Successfully connected to MongoDB!")

	collection := client.Database("taskdb").Collection("tasks")

	// ==========================================
	// PART A: INSERT A DOCUMENT
	// ==========================================
	newTask := models.Task{
		Title:       "Learn MongoDB",
		Description: "Insert my first document using Go",
		DueDate:     "2025-02-25",
		Status:      "In Progress",
	}

	insertResult, err := collection.InsertOne(context.TODO(), newTask)
	if err != nil {
		log.Fatal("Failed to insert task:", err)
	}

	fmt.Println("✅ Inserted Task with ID:", insertResult.InsertedID)

	// ==========================================
	// PART B: READ A DOCUMENT
	// ==========================================

	var foundTask models.Task
	filter := bson.M{"_id": insertResult.InsertedID}
	err = collection.FindOne(context.TODO(), filter).Decode(&foundTask)
	if err != nil {
		log.Fatal("Failed to find task:", err)
	}

	fmt.Printf("🔍 Found Task: %+v\n", foundTask)

	// ==========================================
	// PART C: UPDATE A DOCUMENT
	// ==========================================

	updateRule := bson.M{
		"$set": bson.M{
			"status": "Completed",
		},
	}

	updateRes, err := collection.UpdateOne(context.TODO(), filter, updateRule)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("🔄 Updated %v document(s)\n", updateRes.ModifiedCount)

	// ==========================================
	// PART D: DELETE A DOCUMENT
	// ==========================================

	deleteRes, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("🗑️ Deleted %v document(s)\n", deleteRes.DeletedCount)
}