package analyzer

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
	"loganalyzer/internal/config"
)

type LogResult struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

type Analyzer struct {
	mu sync.Mutex
}

func NewAnalyzer() *Analyzer {
	return &Analyzer{}
}

func (a *Analyzer) AnalyzeLogs(configs []config.LogConfig, statusFilter string) []LogResult {
	var wg sync.WaitGroup
	results := make([]LogResult, 0, len(configs))
	resultsChan := make(chan LogResult, len(configs))

	for _, cfg := range configs {
		wg.Add(1)
		go func(cfg config.LogConfig) {
			defer wg.Done()
			result := a.analyzeLog(cfg)
			if statusFilter == "" || result.Status == statusFilter {
				resultsChan <- result
			}
		}(cfg)
	}

	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	for result := range resultsChan {
		results = append(results, result)
	}

	return results
}

func (a *Analyzer) analyzeLog(cfg config.LogConfig) LogResult {
	time.Sleep(time.Duration(50+rand.Intn(150)) * time.Millisecond)

	result := LogResult{
		ID: cfg.ID,
	}

	if rand.Float32() < 0.8 {
		result.Status = "OK"
	} else {
		result.Status = "FAILED"
		result.Error = fmt.Sprintf("Erreur lors de l'analyse du log %s", cfg.ID)
	}

	return result
}

func (a *Analyzer) PrintResults(results []LogResult) {
	var success, failed int

	fmt.Println("\nRésultats de l'analyse:")
	fmt.Println("------------------------")

	for _, result := range results {
		if result.Status == "OK" {
			success++
			fmt.Printf("✓ %s: OK\n", result.ID)
		} else {
			failed++
			fmt.Printf("✗ %s: FAILED - %s\n", result.ID, result.Error)
		}
	}

	fmt.Printf("\nRésumé:\n")
	fmt.Printf("- Total: %d\n", len(results))
	fmt.Printf("- Succès: %d\n", success)
	fmt.Printf("- Échecs: %d\n", failed)
} 