package repository

import (
	"errors"
	"math/rand"

	"github.com/KaBel11/RandomFact/fact-service/internal/dtos"
	"github.com/KaBel11/RandomFact/fact-service/internal/model"
)

type FactsRepository struct {
	facts []model.Fact
}

func NewFactsRepository() *FactsRepository {
	return &FactsRepository{}
}

func (r *FactsRepository) GetAll() (*[]model.Fact, error) {
	return &r.facts, nil
}

func (r *FactsRepository) GetRandom() (*model.Fact, error) {
	if len(r.facts) == 0 {
		return nil, errors.New("no facts available")
	}
	randomIndex := rand.Intn(len(r.facts))
	return &r.facts[randomIndex], nil
}

func (r *FactsRepository) GetByID(id int) (*model.Fact, error) {
	for i := range r.facts {
		if r.facts[i].ID == id {
			return &r.facts[i], nil
		}
	}
	return nil, errors.New("fact not found")
}

func (r *FactsRepository) Create(parameters dtos.CreateFactRequest) (*model.Fact, error) {
	newFactID := len(r.facts) + 1

	newFact := model.Fact{
		ID:   newFactID,
		Text: parameters.Text,
	}

	r.facts = append(r.facts, newFact)

	return &newFact, nil
}

func (r *FactsRepository) Update(parameters dtos.UpdateFactRequest) (*model.Fact, error) {
	for i := range r.facts {
		if r.facts[i].ID == parameters.ID {
			r.facts[i].Text = parameters.Text
			return &r.facts[i], nil
		}
	}

	return nil, errors.New("fact not found")
}

func (r *FactsRepository) Delete(id int) error {
	for i := range r.facts {
		if r.facts[i].ID == id {
			r.facts = append(r.facts[:i], r.facts[i+1:]...)
			return nil
		}
	}

	return errors.New("fact not found")
}
