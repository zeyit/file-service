package repository

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/zeyd17/file-microservice/model"
)

type IFileRepo interface {
	GetByID(id string) (*model.File, error)
	Create(f *model.File) error
	Delete(id string) error
}

type FileRepo struct {
	db *gorm.DB
}

func NewFileRepo(db *gorm.DB) IFileRepo {
	db.AutoMigrate(&model.File{})
	return &FileRepo{db: db}
}

func (file *FileRepo) GetByID(id string) (*model.File, error) {
	f := &model.File{}
	if err := file.db.Where("Id = ?", id).First(f).Error; err != nil {
		return nil, errors.New("Not fount")
	}
	return f, nil
}
func (file *FileRepo) Create(f *model.File) error {
	return file.db.Create(f).Error
}
func (file *FileRepo) Delete(id string) error {
	f, err := file.GetByID(id)
	if err != nil {
		return err
	}
	if err = file.db.Delete(f).Error; err != nil {
		return err
	}
	return nil
}
