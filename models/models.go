package models

type Paths map[string]Path

type Request struct {
	In          string      `json:"in"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Required    bool        `json:"required"`
	Type        string      `json:"type"`
	IsToModel   bool        `json:"is_to_model"`
	Schema      interface{} `json:"schema"`
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type Path struct {
	Method      string     `json:"method"`
	OperationID string     `json:"operation_id"`
	Summary     string     `json:"summary"`
	Tag         []string   `json:"tag"`
	Consumes    []string   `json:"consumes"`
	Request     []Request  `json:"request"`
	Response    []Response `json:"response"`
}
