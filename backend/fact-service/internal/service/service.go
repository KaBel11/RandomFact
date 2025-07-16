package service

import (
	"github.com/KaBel11/RandomFact/fact-service/internal/dtos"
	"github.com/KaBel11/RandomFact/fact-service/internal/model"
	"github.com/KaBel11/RandomFact/fact-service/internal/repository"
)

type FactsService struct {
	repo *repository.FactsRepository
}

func NewFactsService(repo *repository.FactsRepository) *FactsService {
	return &FactsService{repo: repo}
}

func (s *FactsService) GetAllFacts() (*[]model.Fact, error) {
	return s.repo.GetAll()
}

func (s *FactsService) GetByID(id int) (*model.Fact, error) {
	return s.repo.GetByID(id)
}

func (s *FactsService) GetRandomFact() (*model.Fact, error) {
	return s.repo.GetRandom()
}

func (s *FactsService) CreateFact(parameters dtos.CreateFactRequest) (*model.Fact, error) {
	return s.repo.Create(parameters)
}

func (s *FactsService) UpdateFact(parameters dtos.UpdateFactRequest) (*model.Fact, error) {
	return s.repo.Update(parameters)
}

func (s *FactsService) DeleteFact(id int) error {
	return s.repo.Delete(id)
}
