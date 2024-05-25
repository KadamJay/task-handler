package service

import (
	"context"
	repository "task-handler/internal/repo"
	"task-handler/pkg/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type TaskService struct {
	Repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{
		Repo: repo,
	}
}

func (s *TaskService) GetAllTasks(ctx context.Context) ([]models.Task, error) {
	return s.Repo.GetAllTasks(ctx)
}

func (s *TaskService) CreateTask(ctx context.Context, task models.Task) (*mongo.InsertOneResult, error) {
	return s.Repo.CreateTask(ctx, task)
}
