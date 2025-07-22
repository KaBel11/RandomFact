package repository

import (
	"errors"
	"math/rand"

	"github.com/KaBel11/RandomFact/fact-service/internal/dtos"
	"github.com/KaBel11/RandomFact/fact-service/internal/model"
)

type LocalFactsRepository struct {
	facts []model.Fact
}

func NewLocalFactsRepository() *LocalFactsRepository {
	return &LocalFactsRepository{}
}

func (r *LocalFactsRepository) GetAll() ([]model.Fact, error) {
	return r.facts, nil
}

func (r *LocalFactsRepository) GetRandom() (*model.Fact, error) {
	if len(r.facts) == 0 {
		return nil, errors.New("no facts available")
	}
	randomIndex := rand.Intn(len(r.facts))
	return &r.facts[randomIndex], nil
}

func (r *LocalFactsRepository) GetByID(id uint64) (*model.Fact, error) {
	for i := range r.facts {
		if r.facts[i].ID == id {
			return &r.facts[i], nil
		}
	}
	return nil, errors.New("fact not found")
}

func (r *LocalFactsRepository) Create(parameters dtos.CreateFactRequest) (*model.Fact, error) {
	newFactID := len(r.facts) + 1

	newFact := model.Fact{
		ID:   uint64(newFactID),
		Text: parameters.Text,
	}

	r.facts = append(r.facts, newFact)

	return &newFact, nil
}

func (r *LocalFactsRepository) Update(parameters dtos.UpdateFactRequest) (*model.Fact, error) {
	for i := range r.facts {
		if r.facts[i].ID == parameters.ID {
			r.facts[i].Text = parameters.Text
			return &r.facts[i], nil
		}
	}

	return nil, errors.New("fact not found")
}

func (r *LocalFactsRepository) Delete(id uint64) error {
	for i := range r.facts {
		if r.facts[i].ID == id {
			r.facts = append(r.facts[:i], r.facts[i+1:]...)
			return nil
		}
	}

	return errors.New("fact not found")
}
