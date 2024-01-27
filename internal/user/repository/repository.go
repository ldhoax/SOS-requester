package repository

import (
	"github.com/GoldenOwlAsia/go-golang-api/internal/user/model"
	"github.com/GoldenOwlAsia/go-golang-api/pkg/db"
	"gorm.io/gorm"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{Db: db}
}

func (r Repository) Find(id uint) (model.User, error) {
	user := model.User{}
	result := r.Db.First(&user, id)

	err := result.Error
	return user, db.HandleError(err)
}

func (r Repository) Create(entity *model.User) error {
	var err error
	if err != nil {
		return db.HandleError(err)
	}
	result := r.Db.Create(entity)
	err = result.Error

	if err != nil {
		return db.HandleError(err)
	}
	return db.HandleError(err)
}
