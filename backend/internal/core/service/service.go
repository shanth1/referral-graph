package service

import (
	"github.com/shanth1/graph/internal/core/domain"
	"github.com/shanth1/graph/internal/core/ports"
)

type referralService struct {
	repo ports.GraphRepository
}

func NewReferralService(repo ports.GraphRepository) ports.ReferralService {
	return &referralService{repo: repo}
}

func (s *referralService) GetGraphData() ([]domain.Node, []domain.Edge, error) {
	return s.repo.GetFullGraph()
}
