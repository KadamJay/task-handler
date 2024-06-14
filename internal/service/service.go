package service

import (
	"context"
	"encoding/json"
	"log"
	repository "task-handler/internal/repo"
	"task-handler/pkg/models"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskService struct {
	Repo  *repository.TaskRepository
	Redis *redis.Client
}

func NewTaskService(repo *repository.TaskRepository, redis *redis.Client) *TaskService {
	return &TaskService{
		Repo:  repo,
		Redis: redis,
	}
}

func (s *TaskService) GetAllTasks(ctx context.Context) ([]models.Task, error) {
	var tasks []models.Task

	cachedTasks, err := s.Redis.Get(ctx, "tasks").Result()
	if err == redis.Nil {
		log.Println("cache miss")
		// Cache miss, retrieve tasks from database
		tasks, err = s.Repo.GetAllTasks(ctx)
		if err != nil {
			log.Println("Error retrieving tasks from database:", err)
			return nil, err
		}

		tasksJSON, err := json.Marshal(tasks)
		if err != nil {
			log.Println("Error marshalling tasks:", err)
			return nil, err
		}

		err = s.Redis.Set(ctx, "tasks", tasksJSON, 0).Err()
		if err != nil {
			log.Println("Error setting tasks in Redis:", err)
			return nil, err
		}
	} else if err != nil {
		// Any other error from Redis
		log.Println("Error retrieving tasks from Redis:", err)
		return nil, err
	} else {

		log.Println("cache hit")
		// Cache hit, unmarshal the cached tasks
		err = json.Unmarshal([]byte(cachedTasks), &tasks)
		if err != nil {
			log.Println("Error unmarshalling cached tasks:", err)
			return nil, err
		}
	}

	return tasks, nil
}

func (s *TaskService) CreateTask(ctx context.Context, task models.Task) (*mongo.InsertOneResult, error) {
	result, err := s.Repo.CreateTask(ctx, task)
	if err != nil {
		log.Println("Error creating task:", err)
		return nil, err
	}

	// Invalidate the cache
	err = s.Redis.Del(ctx, "tasks").Err()
	if err != nil {
		log.Printf("Failed to invalidate cache: %v", err)
	}

	return result, nil
}
