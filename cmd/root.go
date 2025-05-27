package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)


var rootCmd = &cobra.Command{
	Use:   "loganalyzer",
	Short: "Un outil d'analyse de logs distribué",
	Long: `LogAnalyzer est un outil en ligne de commande pour analyser des fichiers de logs
de manière distribuée et concurrente.

Il permet de:
- Analyser plusieurs fichiers de logs en parallèle
- Exporter les résultats au format JSON
- Gérer les erreurs de manière robuste
- Filtrer les résultats par statut

Exemple d'utilisation:
  loganalyzer analyze --config config.json --output results.json
  loganalyzer add-log --id web-server --path /var/log/nginx.log --type nginx --file config.json`,
	Version: "1.0.0",
}


func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erreur: %v\n", err)
		os.Exit(1)
	}
} 