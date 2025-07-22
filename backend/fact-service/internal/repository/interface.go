package repository

import (
	"github.com/KaBel11/RandomFact/fact-service/internal/dtos"
	"github.com/KaBel11/RandomFact/fact-service/internal/model"
)

type FactsRepository interface {
	GetAll() ([]model.Fact, error)
    GetRandom() (*model.Fact, error)
	GetByID(id uint64) (*model.Fact, error)
    Create(parameters dtos.CreateFactRequest) (*model.Fact, error)
    Update(parameters dtos.UpdateFactRequest) (*model.Fact, error)
    Delete(id uint64) error
}