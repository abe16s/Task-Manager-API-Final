package repositories

import (
	"context"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/abe16s/Go-Backend-Learning-path/task_manager/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type UserRepository struct {
	collection *mongo.Collection
}

// NewUserRepository creates a new UserRepository.
func NewUserRepository(client *mongo.Client, dbName, collectionName string) *UserRepository {
	collection := client.Database(dbName).Collection(collectionName)

	// check if there is an index on the username field

	// Get a list of existing indexes
    cursor, err := collection.Indexes().List(context.TODO())
    if err != nil {
        log.Printf("could not list indexes: %v", err)
    }
    defer cursor.Close(context.TODO())

    var indexes []bson.M
    if err := cursor.All(context.TODO(), &indexes); err != nil {
        log.Printf("could not parse indexes: %v", err)
    }

    // Check if the "username" index already exists
    indexExists := false
    for _, index := range indexes {
        key := index["key"].(bson.M)
        if len(key) == 1 && key["username"] != nil {
            indexExists = true
            break
        }
    }

    // If the index does not exist, create it
    if !indexExists {
		indexModel := mongo.IndexModel{
			Keys:    bson.D{{Key: "username", Value: 1}}, // Create index on the "username" field
			Options: options.Index().SetUnique(true),    // Ensure the index is unique
		}
		
		// Create the index
		_, err = collection.Indexes().CreateOne(context.TODO(), indexModel)
		if err != nil {
			log.Printf("could not create index: %v", err)
		}
	} else {
		log.Println("username index already exists")
	}
	
	return &UserRepository{
		collection: collection,
	}
}


func (ur *UserRepository) Count() (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	count, err := ur.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, err
	}

	return count, nil
}

// register new user with unique username and password
func (ur *UserRepository) RegisterUser(user *domain.User) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Check if user already exists
	for {
		_, err := ur.collection.InsertOne(ctx, user)
		
		// if user exists return error
		if mongo.IsDuplicateKeyError(err) {
			// Check if the duplicate key error is caused by the username field
			if strings.Contains(err.Error(), "username") {
				return nil, errors.New("username already exists")
			}
			// Check if the duplicate key error is caused by the _id field
			if strings.Contains(err.Error(), "_id") {
				continue
			}
		} else if err != nil {
			return nil, err
		}

		// else create new user
		return user, nil
	}
}


// login user 
func (ur *UserRepository) GetUser(username string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var existingUser domain.User
	err := ur.collection.FindOne(ctx, bson.M{"username": username}).Decode(&existingUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &existingUser, nil
}


// promote user to admin
func (ur *UserRepository) PromoteUser(username string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	result, err := ur.collection.UpdateOne(ctx, bson.D{{Key: "username", Value: username}}, bson.D{{Key: "$set", Value: bson.M{"is_admin": true}}})
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("username not found")
	} 
	
	return nil
}