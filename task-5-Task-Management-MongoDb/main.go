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
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// 1. Point to your local database
	uri := "mongodb://localhost:27017"

	// 2. Connect
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	// 3. Ping to verify connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Connection failed! Is MongoDB running? Error:", err)
	}

	fmt.Println("✅ Successfully connected to MongoDB!")
    
    // 4. Close the connection when done
    client.Disconnect(context.TODO())
}