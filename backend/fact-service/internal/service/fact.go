package service

import (
	"context"

	"github.com/KaBel11/RandomFact/fact-service/internal/dtos"
	"github.com/KaBel11/RandomFact/fact-service/internal/model"
	"github.com/KaBel11/RandomFact/fact-service/internal/repository"
)

type FactService struct {
	repo *repository.FactRepository
}

func NewFactService(repo *repository.FactRepository) *FactService {
	return &FactService{repo: repo}
}

func (s *FactService) GetAllFacts(ctx context.Context) ([]model.Fact, error) {
	return s.repo.GetAll(ctx)
}

func (s *FactService) GetByID(ctx context.Context, id uint64) (*model.Fact, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *FactService) GetRandomFact(ctx context.Context) (*model.Fact, error) {
	return s.repo.GetRandom(ctx)
}

func (s *FactService) CreateFact(ctx context.Context, parameters dtos.CreateFactRequest) (*model.Fact, error) {
	return s.repo.Create(ctx, parameters)
}

func (s *FactService) UpdateFact(ctx context.Context, parameters dtos.UpdateFactRequest) (*model.Fact, error) {
	return s.repo.Update(ctx, parameters)
}

func (s *FactService) DeleteFact(ctx context.Context, id uint64) error {
	return s.repo.Delete(ctx, id)
}