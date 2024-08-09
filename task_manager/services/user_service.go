package services

import (
	"context"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/abe16s/Go-Backend-Learning-path/task_manager/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


var jwtSecret = []byte("qwertyuiopasdfghjklzxcvbnm")

type UserService struct {
	collection *mongo.Collection
}

// NewUserService creates a new UserService.
func NewUserService(client *mongo.Client, dbName, collectionName string) *UserService {
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
	
	return &UserService{
		collection: collection,
	}
}


// register new user with unique username and password
func (s *UserService) RegisterUser(user models.User) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
		return nil, err
    }

    user.Password = string(hashedPassword)

	// check if user is empty in database and it is the first user if it is make isAdmin = true
	count, err := s.collection.CountDocuments(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    if count == 0 {
        user.IsAdmin = true
    }

	// Check if user already exists
	for {
		user.ID = uuid.New()
		_, err := s.collection.InsertOne(ctx, user)
		
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
		return &user, nil
	}
}


// login user 
func (s *UserService) LoginUser(user models.User) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var existingUser models.User
	err := s.collection.FindOne(ctx, bson.M{"username": user.Username}).Decode(&existingUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", errors.New("user not found")
		}
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": existingUser.Username,
		"is_admin": existingUser.IsAdmin,
		"exp": time.Now().Add(900),
	})
	
	jwtToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", errors.New("internal server error")
	}

	return jwtToken, nil
}


// promote user to admin
func (s *UserService) PromoteUser(username string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := s.collection.UpdateOne(ctx, bson.D{{Key: "username", Value: username}}, bson.D{{Key: "$set", Value: bson.M{"is_admin": true}}})
	if err != nil {
		return err
	}

	return nil
}