package entity

type Request struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}

type Response struct {
	Embeddings []float64 `json:"embedding"`
}
