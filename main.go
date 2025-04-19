package main

import (
	"go-gpt/internal/config"
	"go-gpt/internal/llm"
	"go-gpt/internal/ui"
	"log"
)

func main() {
	cfg, err := config.LoadEnv()
	if err != nil {
		log.Fatalf("unable to get config: %v\n", err)
	}

	svc := llm.InitLLMService(*cfg)
	llm := llm.InitLLM(svc)

	app := ui.NewApp(llm)
	if err := app.StartServer(); err != nil {
		log.Fatal(err)
	}
}
