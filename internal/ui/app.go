package ui

import (
	"context"
	"embed"
	"fmt"
	"go-gpt/internal/llm"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

// App struct
type App struct {
	llm *llm.LLM
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp(llm *llm.LLM) *App {
	return &App{llm: llm}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	fmt.Println("name received: ", name)
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) Prompt(prompt string) string {
	fmt.Println("prompt received: ", prompt)
	// TODO: return the response back to UI
	res, err := a.llm.Get(prompt)
	if err != nil {
		// TODO: handle proper error response
		log.Fatal(err)
	}
	// fmt.Println("response received: ", res)
	if res != "" {
		return res
	}
	return ""
}

func (a *App) StartServer() error {
	// Create application with options
	err := wails.Run(&options.App{
		Title:  "go-gpt",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        a.startup,
		Bind: []interface{}{
			a,
		},
	})

	if err != nil {
		return fmt.Errorf("unable to start the app: %v", err)
	}
	return nil
}
