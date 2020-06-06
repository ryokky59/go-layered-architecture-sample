package infra

import (
	"sample/domain/model"
	"sample/domain/repository"

	"github.com/jinzhu/gorm"
)

type taskRepository struct {
	Conn *gorm.DB
}

// NewTaskRepository task repositoryのコンストラクタ
func NewTaskRepository(conn *gorm.DB) repository.TaskRepository {
	return &taskRepository{Conn: conn}
}

// Create taskの保存
func (tr *taskRepository) Create(task *model.Task) (*model.Task, error) {
	if err := tr.Conn.Create(&task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

// FindByID taskをIDで取得
func (tr *taskRepository) FindByID(id int) (*model.Task, error) {
	task := &model.Task{ID: id}

	if err := tr.Conn.First(&task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

// Update taskの更新
func (tr *taskRepository) Update(task *model.Task) (*model.Task, error) {
	if err := tr.Conn.Model(&task).Update(&task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

// Delete taskの削除
func (tr *taskRepository) Delete(task *model.Task) error {
	if err := tr.Conn.Delete(&task).Error; err != nil {
		return err
	}

	return nil
}
