package usecase

import (
	"sample/domain/model"
	"sample/domain/repository"
)

type TaskUsecase interface {
	Create(task *model.Task) (*model.Task, error)
	FindByID(id int) (*model.Task, error)
	Update(id int, task *model.Task) (*model.Task, error)
	Delete(id int) error
}

type taskUsecase struct {
	taskRepo repository.TaskRepository
}

func NewTaskUsecase(taskRepo repository.TaskRepository) TaskUsecase {
	return &taskUsecase{taskRepo: taskRepo}
}

func (tu *taskUsecase) Create(task *model.Task) (*model.Task, error) {
	createdTask, err := tu.taskRepo.Create(task)
	if err != nil {
		return nil, err
	}

	return createdTask, nil
}

func (tu *taskUsecase) FindByID(id int) (*model.Task, error) {
	foundTask, err := tu.taskRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return foundTask, nil
}

func (tu *taskUsecase) Update(id int, task *model.Task) (*model.Task, error) {
	task.ID = id
	updatedTask, err := tu.taskRepo.Update(task)
	if err != nil {
		return nil, err
	}

	return updatedTask, nil
}

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
