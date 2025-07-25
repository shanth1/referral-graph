package graphdb

import (
	"fmt"
	"strconv"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/shanth1/graph/internal/core/domain"
)

type neo4jRepository struct {
	driver neo4j.Driver
}

func NewNeo4jRepository(uri, user, password string) (*neo4jRepository, error) {
	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(user, password, ""))
	if err != nil {
		return nil, err
	}
	return &neo4jRepository{driver: driver}, nil
}

func (r *neo4jRepository) GetFullGraph() ([]domain.Node, []domain.Edge, error) {
	session := r.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()

	// Используем APOC для удобного экспорта в JSON-подобную структуру
	cypherQuery := `
        MATCH (n)
        OPTIONAL MATCH (n)-[r]->(m)
        RETURN n, r, m
    `
	nodes := make(map[string]domain.Node)
	var edges []domain.Edge

	_, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		result, err := tx.Run(cypherQuery, nil)
		if err != nil {
			return nil, err
		}

		for result.Next() {
			record := result.Record()

			// Обработка узлов
			if nodeVal, ok := record.Get("n"); ok && nodeVal != nil {
				node := nodeVal.(neo4j.Node)
				if _, exists := nodes[strconv.FormatInt(node.Id, 10)]; !exists {
					nodes[strconv.FormatInt(node.Id, 10)] = domain.Node{
						ID:    strconv.FormatInt(node.Id, 10),
						Label: node.Props["name"].(string),
					}
				}
			}
			if nodeVal, ok := record.Get("m"); ok && nodeVal != nil {
				node := nodeVal.(neo4j.Node)
				if _, exists := nodes[strconv.FormatInt(node.Id, 10)]; !exists {
					nodes[strconv.FormatInt(node.Id, 10)] = domain.Node{
						ID:    strconv.FormatInt(node.Id, 10),
						Label: node.Props["name"].(string),
					}
				}
			}

			// Обработка связей
			if relVal, ok := record.Get("r"); ok && relVal != nil {
				rel := relVal.(neo4j.Relationship)
				edges = append(edges, domain.Edge{
					ID:     strconv.FormatInt(rel.Id, 10),
					Source: strconv.FormatInt(rel.StartId, 10),
					Target: strconv.FormatInt(rel.EndId, 10),
				})
			}
		}
		return nil, result.Err()
	})

	if err != nil {
		return nil, nil, err
	}

	// Преобразование map в slice
	nodeSlice := make([]domain.Node, 0, len(nodes))
	for _, node := range nodes {
		nodeSlice = append(nodeSlice, node)
	}

	// Если нет данных, создадим тестовые
	if len(nodeSlice) == 0 {
		return r.createInitialData()
	}

	return nodeSlice, edges, nil
}

// createInitialData создает начальные данные, если база пуста
func (r *neo4jRepository) createInitialData() ([]domain.Node, []domain.Edge, error) {
	session := r.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		cypher := `
            CREATE (u1:User {name: 'Alice'}),
                   (u2:User {name: 'Bob'}),
                   (u3:User {name: 'Charlie'}),
                   (u4:User {name: 'David'}),
                   (u1)-[:INVITED]->(u2),
                   (u1)-[:INVITED]->(u3),
                   (u2)-[:INVITED]->(u4)
        `
		return tx.Run(cypher, nil)
	})

	if err != nil {
		return nil, nil, fmt.Errorf("could not create initial data: %w", err)
	}

	// После создания данных, запрашиваем их снова
	return r.GetFullGraph()
}
