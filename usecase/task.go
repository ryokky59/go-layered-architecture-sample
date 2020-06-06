package usecase

import (
	"sample/domain/model"
	"sample/domain/repository"
)

// TaskUsecase task usecaseのinterface
type TaskUsecase interface {
	Create(title, content string) (*model.Task, error)
	FindByID(id int) (*model.Task, error)
	Update(id int, title, content string) (*model.Task, error)
	Delete(id int) error
}

type taskUsecase struct {
	taskRepo repository.TaskRepository
}

// NewTaskUsecase task usecaseのコンストラクタ
func NewTaskUsecase(taskRepo repository.TaskRepository) TaskUsecase {
	return &taskUsecase{taskRepo: taskRepo}
}

// Create taskを保存するときのユースケース
func (tu *taskUsecase) Create(title, content string) (*model.Task, error) {
	task, err := model.NewTask(title, content)
	if err != nil {
		return nil, err
	}

	createdTask, err := tu.taskRepo.Create(task)
	if err != nil {
		return nil, err
	}

	return createdTask, nil
}

// FindByID taskをIDで取得するときのユースケース
func (tu *taskUsecase) FindByID(id int) (*model.Task, error) {
	foundTask, err := tu.taskRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return foundTask, nil
}

// Update taskを更新するときのユースケース
func (tu *taskUsecase) Update(id int, title, content string) (*model.Task, error) {
	targetTask, err := tu.taskRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	err = targetTask.Set(title, content)
	if err != nil {
		return nil, err
	}

	updatedTask, err := tu.taskRepo.Update(targetTask)
	if err != nil {
		return nil, err
	}

	return updatedTask, nil
}

// Delete taskを削除するときのユースケース
func (tu *taskUsecase) Delete(id int) error {
	task, err := tu.taskRepo.FindByID(id)
	if err != nil {
		return err
	}

	err = tu.taskRepo.Delete(task)
	if err != nil {
		return err
	}

	return nil
}
