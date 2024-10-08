package repository

import (
	"github.com/GoldenOwlAsia/go-golang-api/internal/request/model"
	"github.com/GoldenOwlAsia/go-golang-api/pkg/db"
	"gorm.io/gorm"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{Db: db}
}

func (r Repository) FindAll() ([]model.Request, error) {
	var requests []model.Request
	result := r.Db.Find(&requests)
	return requests, db.HandleError(result.Error)
}

func (r Repository) Find(id uint) (model.Request, error) {
	request := model.Request{}
	result := r.Db.First(&request, id)
	return request, db.HandleError(result.Error)
}

func (r Repository) Create(entity *model.Request) error {
	return db.HandleError(r.Db.Create(entity).Error)
}

func (r Repository) Update(entity *model.Request) error {
	return db.HandleError(r.Db.Save(entity).Error)
}

func (r Repository) Delete(id uint) error {
	return db.HandleError(r.Db.Delete(&model.Request{}, id).Error)
}