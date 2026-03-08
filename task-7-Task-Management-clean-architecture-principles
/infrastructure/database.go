package infrastructure

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Database {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Database ping failed:", err)
	}

	log.Println("✅ Database connected successfully")
	db := client.Database("taskdb")

	userCollection := db.Collection("users")
	
	indexModel := mongo.IndexModel{
		Keys: bson.D{{Key: "username", Value: 1}}, 
		Options: options.Index().SetUnique(true),
	}

	_, err = userCollection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		log.Println("⚠️ Warning: Could not create username index:", err)
	}
	log.Println("✅ Unique Index on 'username' is ready!")
	

	return db
}