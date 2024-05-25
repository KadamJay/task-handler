package repository

import (
	"context"

	"task-handler/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository struct {
	Collection *mongo.Collection
}

func NewTaskRepository(collection *mongo.Collection) *TaskRepository {
	return &TaskRepository{
		Collection: collection,
	}
}
func (r *TaskRepository) GetAllTasks(ctx context.Context) ([]models.Task, error) {
	var tasks []models.Task
	// ASK
	cursor, err := r.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	// ASK
	if err = cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *TaskRepository) CreateTask(ctx context.Context, task models.Task) (*mongo.InsertOneResult, error) {
	// ASK
	return r.Collection.InsertOne(ctx, task)
}
