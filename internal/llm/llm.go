package llm

func InitLLM(svc LLMService) *LLM {
	return &LLM{
		LLMService: svc,
	}
}

// client
type LLM struct {
	LLMService LLMService
}

func (llm LLM) Get(prompt string) (string, error) {
	res, err := llm.LLMService.Get(prompt)
	if err != nil {
		return "", err
	}
	return res, nil
}
