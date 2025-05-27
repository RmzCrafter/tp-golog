package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"loganalyzer/internal/analyzer"
	"loganalyzer/internal/config"
	"loganalyzer/internal/reporter"
)

var (
	configPath    string
	outputPath    string
	statusFilter  string
	addTimestamp  bool
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyse les fichiers de logs spécifiés dans la configuration",
	Long: `La commande analyze lit un fichier de configuration JSON contenant
la liste des logs à analyser et les traite de manière concurrente.

Chaque log est analysé dans une goroutine séparée pour optimiser les performances.
Les résultats peuvent être filtrés par statut et exportés au format JSON.

Exemples:
  # Analyse de base avec configuration
  loganalyzer analyze --config config.json

  # Analyse avec export des résultats
  loganalyzer analyze -c config.json -o results.json

  # Analyse avec filtre de statut
  loganalyzer analyze -c config.json --status FAILED

  # Analyse avec horodatage automatique
  loganalyzer analyze -c config.json -o results.json --timestamp`,
	Run: func(cmd *cobra.Command, args []string) {
		runAnalyze()
	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)

	
	analyzeCmd.Flags().StringVarP(&configPath, "config", "c", "", "Chemin vers le fichier de configuration JSON (requis)")
	analyzeCmd.Flags().StringVarP(&outputPath, "output", "o", "", "Chemin vers le fichier de sortie JSON (optionnel)")
	analyzeCmd.Flags().StringVar(&statusFilter, "status", "", "Filtrer les résultats par statut (OK ou FAILED)")
	analyzeCmd.Flags().BoolVar(&addTimestamp, "timestamp", false, "Ajouter un horodatage au nom du fichier de sortie")

	
	analyzeCmd.MarkFlagRequired("config")
}

func runAnalyze() {
	
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Erreur: Le fichier de configuration '%s' n'existe pas.\n", configPath)
		os.Exit(1)
	}

	
	fmt.Printf("Chargement de la configuration depuis: %s\n", configPath)
	configs, err := config.LoadConfig(configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erreur lors du chargement de la configuration: %v\n", err)
		os.Exit(1)
	}

	if len(configs) == 0 {
		fmt.Println("Aucun log à analyser dans la configuration.")
		return
	}

	fmt.Printf("Nombre de logs à analyser: %d\n", len(configs))

	
	if statusFilter != "" && statusFilter != "OK" && statusFilter != "FAILED" {
		fmt.Fprintf(os.Stderr, "Erreur: Le statut '%s' n'est pas valide. Utilisez 'OK' ou 'FAILED'.\n", statusFilter)
		os.Exit(1)
	}

	
	analyzer := analyzer.NewAnalyzer()
	fmt.Println("\nDémarrage de l'analyse concurrente...")
	
	results := analyzer.AnalyzeLogs(configs, statusFilter)

	
	analyzer.PrintResults(results)

	
	if outputPath != "" {
		reporter := reporter.NewReporter()
		
	
		if err := reporter.ValidateOutputPath(outputPath); err != nil {
			fmt.Fprintf(os.Stderr, "Erreur de validation du chemin de sortie: %v\n", err)
			os.Exit(1)
		}

		
		if err := reporter.ExportResults(results, outputPath, addTimestamp); err != nil {
			fmt.Fprintf(os.Stderr, "Erreur lors de l'export: %v\n", err)
			os.Exit(1)
		}
	}

	if statusFilter != "" {
		fmt.Printf("\nAnalyse terminée. Résultats filtrés par statut: %s\n", statusFilter)
	} else {
		fmt.Println("\nAnalyse terminée.")
	}
} 