package repository

import (
	"context"
	"errors"

	"github.com/philipos/api/domain" 
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)



type userModel struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
	Role     string             `bson:"role"`
}

func fromUserDomain(u *domain.User) userModel {
	objID, _ := primitive.ObjectIDFromHex(u.ID)
	
	return userModel{
		ID:       objID,
		Username: u.Username,
		Password: u.Password,
		Role:     u.Role,
	}
}

func toUserDomain(m userModel) domain.User {
	return domain.User{
		ID:       m.ID.Hex(),
		Username: m.Username,
		Password: m.Password,
		Role:     m.Role,
	}
}

type userRepository struct {
	database   *mongo.Database
	collection string
}

func NewUserRepository(db *mongo.Database, collection string) domain.UserRepository {
	return &userRepository{
		database:   db,
		collection: collection,
	}
}

func (r *userRepository) Create(user *domain.User) error {
	collection := r.database.Collection(r.collection)

	dbUser := fromUserDomain(user)

	res, err := collection.InsertOne(context.TODO(), dbUser)
	if err != nil {
		return err
	}

	user.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return nil
}

func (r *userRepository) GetByUsername(username string) (*domain.User, error) {
	collection := r.database.Collection(r.collection)
	var dbUser userModel

	err := collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&dbUser)
	
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("user not found")
	} else if err != nil {
		return nil, err
	}

	domainUser := toUserDomain(dbUser)
	return &domainUser, nil
}