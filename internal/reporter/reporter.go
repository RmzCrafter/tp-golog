package reporter

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
	"loganalyzer/internal/analyzer"
)


type Reporter struct{}


func NewReporter() *Reporter {
	return &Reporter{}
}


func (r *Reporter) ExportResults(results []analyzer.LogResult, outputPath string, addTimestamp bool) error {
	if addTimestamp {
		outputPath = r.addTimestampToFilename(outputPath)
	}

	
	if err := r.createDirectories(outputPath); err != nil {
		return fmt.Errorf("impossible de créer les répertoires pour %s: %w", outputPath, err)
	}


	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("impossible de créer le fichier de sortie %s: %w", outputPath, err)
	}
	defer file.Close()


	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(results); err != nil {
		return fmt.Errorf("erreur lors de l'encodage JSON vers %s: %w", outputPath, err)
	}

	fmt.Printf("Résultats exportés vers: %s\n", outputPath)
	return nil
}


func (r *Reporter) addTimestampToFilename(filename string) string {
	now := time.Now()
	timestamp := now.Format("060102") // Format AAMMJJ (AnnéeMoisJour)
	
	dir := filepath.Dir(filename)
	base := filepath.Base(filename)
	ext := filepath.Ext(base)
	nameWithoutExt := strings.TrimSuffix(base, ext)
	
	newFilename := fmt.Sprintf("%s_%s%s", timestamp, nameWithoutExt, ext)
	return filepath.Join(dir, newFilename)
}


func (r *Reporter) createDirectories(filePath string) error {
	dir := filepath.Dir(filePath)
	if dir == "." || dir == "/" {
		return nil
	}
	
	return os.MkdirAll(dir, 0755)
}

func (r *Reporter) ValidateOutputPath(path string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("impossible de créer le répertoire %s: %w", dir, err)
	}
	return nil
} 