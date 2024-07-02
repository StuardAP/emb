package http

import (
	"net/http"
	"io"
	"encoding/json"
	"bytes"
	"github.com/StuardAP/emb/pkg/domain/entity"
)

type HTTPClient struct {
	BaseURL string
}

func NewHTTPClient(baseURL string) *HTTPClient {
	return &HTTPClient{
		BaseURL: baseURL,
	}
}

func (c *HTTPClient) PostEmbeddings(text string) ([]float64, error) {
	url := c.BaseURL + "/api/embeddings"

	requestBody, err := json.Marshal(entity.Request{
		Model:  "nomic-embed-text",
		Prompt: text,
	})
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
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
