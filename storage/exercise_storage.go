package storage

import (
	"github.com/Csejersen/fitnessTracker/models"
	"gorm.io/gorm"
)

type ExerciseStore interface {
	CreateExercise(*models.Exercise) error
}

type GormExerciseStore struct {
	DB *gorm.DB
}

func NewGormExerciseStore(db *gorm.DB) *GormExerciseStore {
	return &GormExerciseStore{
		DB: db,
	}
}

func (s *GormExerciseStore) CreateExercise(exercise *models.Exercise) error {
	return s.DB.Create(exercise).Error
}
