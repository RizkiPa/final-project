package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"
	"errors"

	"gorm.io/gorm"
)

type TaskRepository interface {
	GetTasks(ctx context.Context, id int) ([]entity.Task, error)
	StoreTask(ctx context.Context, task *entity.Task) (taskId int, err error)
	GetTaskByID(ctx context.Context, id int) (entity.Task, error)
	GetTasksByCategoryID(ctx context.Context, catId int) ([]entity.Task, error)
	UpdateTask(ctx context.Context, task *entity.Task) error
	DeleteTask(ctx context.Context, id int) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) GetTasks(ctx context.Context, id int) ([]entity.Task, error) {
	result := []entity.Task{}
	err := r.db.WithContext(ctx).Model(&entity.Task{}).Where("user_id = ?", id).Find(&result).Error
	if err != nil {
		return []entity.Task{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []entity.Task{}, nil
	}
	return result, nil
}

// TODO: replace this

func (r *taskRepository) StoreTask(ctx context.Context, task *entity.Task) (taskId int, err error) {
	if err := r.db.WithContext(ctx).Create(&task).Error; err != nil {
		return 0, err
	}
	return task.ID, nil // TODO: replace this
}

func (r *taskRepository) GetTaskByID(ctx context.Context, id int) (entity.Task, error) {
	result := entity.Task{}
	err := r.db.WithContext(ctx).Create(&entity.Task{}).Where("id = ?", id).Find(&result).Error
	if err != nil {
		return entity.Task{}, nil

	}
	return result, nil // TODO: replace this
}

func (r *taskRepository) GetTasksByCategoryID(ctx context.Context, catId int) ([]entity.Task, error) {
	result := []entity.Task{}
	if err := r.db.WithContext(ctx).Model(&entity.Task{}).Where("category_id = ?", catId).Find(&result).Error; err != nil {
		return []entity.Task{}, err
	}
	return result, nil // TODO: replace this
}

func (r *taskRepository) UpdateTask(ctx context.Context, task *entity.Task) error {
	if err := r.db.WithContext(ctx).Model(&entity.Task{}).Where("id = ?", task.ID).Updates(&task).Error; err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (r *taskRepository) DeleteTask(ctx context.Context, id int) error {
	if err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&entity.Task{}).Error; err != nil {
		return err
	}
	return nil // TODO: replace this
}
