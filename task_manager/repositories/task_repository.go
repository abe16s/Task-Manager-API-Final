package repositories

import (
	"context"
	"errors"
	"time"
	"github.com/abe16s/Go-Backend-Learning-path/task_manager/domain"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskRepository struct {
	collection *mongo.Collection
}

// NewTaskRepository creates a new TaskRepository.
func NewTaskRepository(client *mongo.Client, dbName, collectionName string) *TaskRepository {
	collection := client.Database(dbName).Collection(collectionName)
	return &TaskRepository{
		collection: collection,
	}
}

func (tr *TaskRepository) GetTasks() ([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
  
	cursor, err := tr.collection.Find(ctx, bson.D{{}})
	if err != nil {
  
	  return nil, err
	}
  
	defer cursor.Close(ctx)
  
	tasks := make([]domain.Task,0)
	for cursor.Next(ctx) {
	  var task domain.Task
	  if err := cursor.Decode(&task); err != nil {
  
		return nil, err
	  }
	  tasks = append(tasks, task)
	}
  
	if err := cursor.Err(); err != nil {
  
	  return nil, err
	}
  
	return tasks, nil
}

func (tr *TaskRepository) GetTaskById(id uuid.UUID) (*domain.Task, error) {
	
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
  
	filter := bson.D{{Key: "_id", Value: id}}
  
	// Find a single document that matches the filter
	var task domain.Task
	err := tr.collection.FindOne(ctx, filter).Decode(&task)
	if err != nil {
	  if err == mongo.ErrNoDocuments {
		return nil, errors.New("task Not Found")
	  }
	  return nil, err
	}
  
	return &task, nil
}

func (tr *TaskRepository) UpdateTaskByID(id uuid.UUID, updatedTask domain.Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
  
	filter := bson.D{{Key: "_id", Value: id}}
  
	update := bson.D{
	  {Key: "$set", Value: bson.D{
		{Key: "title", Value: updatedTask.Title},
		{Key: "description", Value: updatedTask.Description},
		{Key: "due_date", Value: updatedTask.DueDate},
		{Key: "status", Value: updatedTask.Status},
	  }},
	}

	// Update the document that matches the filter
	result :=  tr.collection.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After))
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return errors.New("task not found")
		}
		return result.Err()
	}
  
	return nil
}

func (tr *TaskRepository) DeleteTask(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
    
	filter := bson.D{{Key: "_id", Value: id}}
  
	// Delete the document that matches the filter
	result, err := tr.collection.DeleteOne(ctx, filter)
	if err != nil {
	  return err
	}

	if result.DeletedCount == 0 {
		return errors.New("task not found")
	}

	return nil
}

func (tr *TaskRepository) AddTask(task domain.Task) (*domain.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Recreate task until the ID conflict is resolved
	for {
		task.ID = uuid.New()
  
		_, err := tr.collection.InsertOne(ctx, task)
		if mongo.IsDuplicateKeyError(err) {
			// If a duplicate key error occurs, generate a new ID and try again
			continue
		} else if err != nil {
			return nil, err
		}
		return &task, nil
	}
}