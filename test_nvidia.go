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
	fmt.Printf("Config Loaded:\n - LLM: %s\n - APIKey: %s...\n - APIBase: %s\n",
		cfg.LLM,
		cfg.APIKey[:8],
		cfg.APIBase,
	)

	fmt.Println("\n=== Initializing LLM Client ===")
	// Note: We need a composite resolver or legacy resolver. Since we're using
	// legacy resolver via env vars (XALGORIX_LLM, XALGORIX_API_KEY, XALGORIX_API_BASE),
	// we initialize with WithResolver to use the legacy/composite resolver.
	resolver := llm.NewCompositeResolver(
		llm.WithLegacy(cfg),
	)
	client := llm.NewClient(cfg, llm.WithResolver(resolver))

	fmt.Println("\n=== Testing Chat request ===")
	messages := []llm.Message{
		{Role: "user", Content: "Hello! Tell me in 1 sentence which model you are, and print a greeting."},
	}
	resp, err := client.Chat(messages)
	if err != nil {
		log.Fatalf("Chat failed: %v", err)
	}

	fmt.Printf("\nResponse from model:\n%s\n", resp)
}
