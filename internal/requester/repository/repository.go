package repository

import (
	"github.com/ldhoax/SOS-requester/internal/requester/model"
	"github.com/ldhoax/SOS-requester/pkg/db"
	"gorm.io/gorm"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{Db: db}
}

func (r Repository) Find(id string) (model.Requester, error) {
	requester := model.Requester{}
	result := r.Db.First(&requester, "id = ?", id)
	return requester, db.HandleError(result.Error)
}

func (r Repository) Create(entity *model.Requester) error {
	return db.HandleError(r.Db.Create(entity).Error)
}
