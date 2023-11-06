package models

type Paths map[string]Path

type Request struct {
	In       string      `json:"in"`
	Required bool        `json:"required"`
	Schema   interface{} `json:"schema"`
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
	Request     []Request  `json:"request"`
	Response    []Response `json:"response"`
}
