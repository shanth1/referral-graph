package ports

import "github.com/shanth1/graph/internal/core/domain"

// GraphRepository - это порт для взаимодействия с хранилищем графа
type GraphRepository interface {
	GetFullGraph() ([]domain.Node, []domain.Edge, error)
}

// ReferralService - это порт для бизнес-логики
type ReferralService interface {
	GetGraphData() ([]domain.Node, []domain.Edge, error)
}
