package main

type APIS map[string]APIConfig

type Request struct {
	In  string `json:"in"`
	Raw string `json:"raw"`
}
type Response struct {
	Raw string `json:"raw"`
}
type APIConfig struct {
	Method   string     `json:"method"`
	Request  []Request  `json:"request"`
	Response []Response `json:"response"`
}
