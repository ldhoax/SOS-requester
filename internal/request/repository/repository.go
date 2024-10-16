package repository

import (
	"github.com/ldhoax/SOS-requester/internal/request/model"
	"github.com/ldhoax/SOS-requester/pkg/db"
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

func (r Repository) Find(id string) (model.Request, error) {
	request := model.Request{}
	result := r.Db.First(&request, "id = ?", id)
	return request, db.HandleError(result.Error)
}

func (r Repository) Create(entity *model.Request) error {
	return db.HandleError(r.Db.Create(entity).Error)
}
