package domain

type Node struct {
	ID    string `json:"id"`
	Label string `json:"label"`
}

type Edge struct {
	ID     string `json:"id"`
	Source string `json:"source"`
	Target string `json:"target"`
}
