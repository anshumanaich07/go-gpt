package llm

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"go-gpt/internal/config"
	"go-gpt/internal/domain"
	"log"
	"net/http"
)

type llmSvc struct {
	mime    string
	url     string
	model   string
	httpSvc *http.Client
}

func InitLLMService(cfg config.Config) LLMService {
	return llmSvc{
		httpSvc: &http.Client{},
		url:     cfg.Localhost,
		model:   cfg.Model,
		mime:    "application/json",
	}
}

type LLMService interface {
	Get(query string) (string, error)
}

func (svc llmSvc) Get(prompt string) (string, error) {
	// ollama request
	req := domain.Request{
		Model:  svc.model,
		Prompt: prompt,
		Stream: true,
	}

	// to bytes
	bVal, err := json.Marshal(req)
	if err != nil {
		log.Printf("Error while marshalling: %v", err)
		return "", fmt.Errorf("unable to marshal ollama request: %v", err)
	}

	// post request
	resp, err := svc.httpSvc.Post(svc.url, svc.mime, bytes.NewBuffer(bVal))
	if err != nil {
		log.Printf("error while sending the post request: %v", err)
		return "", err
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	model_resp := ""

	for scanner.Scan() {
		response := domain.Response{}
		if err = json.Unmarshal(scanner.Bytes(), &response); err != nil {
			log.Printf("error while unmarshalling into the Response type: %v", err)
			return "", err
		}
		if response.Done {
			break
		}
		model_resp += response.Response
	}

	return model_resp, nil
}
