package graphdb

import "github.com/shanth1/graph/internal/core/domain"

type mockRepository struct{}

func NewMockRepository() *mockRepository {
	return &mockRepository{}
}

func (m *mockRepository) GetFullGraph() ([]domain.Node, []domain.Edge, error) {
	nodes := []domain.Node{
		{ID: "1", Label: "Mock Alice"},
		{ID: "2", Label: "Mock Bob"},
		{ID: "3", Label: "Mock Charlie"},
	}
	edges := []domain.Edge{
		{ID: "e1", Source: "1", Target: "2"},
		{ID: "e2", Source: "1", Target: "3"},
	}
	return nodes, edges, nil
}
