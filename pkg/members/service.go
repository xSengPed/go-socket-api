package members

import (
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) Create() error {
	tx := s.db.Create(&Member{
		Name: "John",
		Age:  20,
	})

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
