//go:build ignore

package main

import (
	"fmt"
	"log"

	"github.com/xalgord/xalgorix/v4/internal/config"
	"github.com/xalgord/xalgorix/v4/internal/llm"
)

func main() {
	fmt.Println("=== Loading Config ===")
	cfg := config.Get()

	fmt.Println("\n=== Initializing LLM Client ===")
	resolver := llm.NewCompositeResolver(
		llm.WithLegacy(cfg),
	)
	client := llm.NewClient(cfg, llm.WithResolver(resolver))

	fmt.Println("\n=== Testing ChatStream ===")
	messages := []llm.Message{
		{Role: "user", Content: "Hello! Say three words."},
	}
	ch := client.ChatStream(messages)
	for chunk := range ch {
		if chunk.Err != nil {
			log.Fatalf("Stream chunk error: %v", chunk.Err)
		}
		if chunk.Done {
			fmt.Println("\n[Done]")
			break
		}
		fmt.Print(chunk.Content)
	}
}
