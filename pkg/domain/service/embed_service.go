package service

import (
	"encoding/json"
	"io"
	net "net/http"
	"bytes"
	"github.com/StuardAP/emb/pkg/domain/entity"
	"github.com/StuardAP/emb/pkg/infrastructure/http"
)

type EmbedService struct {
	HTTPClient *http.HTTPClient
}

func NewEmbedService(httpClient *http.HTTPClient) *EmbedService {
	return &EmbedService{
		HTTPClient: httpClient,
	}
}

func (s *EmbedService) GetEmbeddings(text string) ([]float64, error) {
	url := s.HTTPClient.BaseURL + "/api/embeddings"

	requestBody, err := json.Marshal(entity.Request{
		Model:  "nomic-embed-text",
		Prompt: text,
	})
	if err != nil {
		return nil, err
	}

	resp, err := net.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var embedResp entity.Response
	if err := json.Unmarshal(body, &embedResp); err != nil {
		return nil, err
	}

	return embedResp.Embeddings, nil
}

