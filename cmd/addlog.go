package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"loganalyzer/internal/config"
)

var (
	logID      string
	logPath    string
	logType    string
	configFile string
)

var addLogCmd = &cobra.Command{
	Use:   "add-log",
	Short: "Ajoute un nouveau log à la configuration",
	Long: `La commande add-log permet d'ajouter un nouveau fichier de log
à la configuration JSON. Si le fichier de configuration n'existe pas,
il sera créé automatiquement.

Exemples:
  # Ajouter un log à une configuration existante
  loganalyzer add-log --id web-server --path /var/log/nginx.log --type nginx --file config.json

  # Créer une nouvelle configuration avec un log
  loganalyzer add-log --id app-log --path /var/log/app.log --type application --file new_config.json`,
	Run: func(cmd *cobra.Command, args []string) {
		runAddLog()
	},
}

func init() {
	rootCmd.AddCommand(addLogCmd)

	addLogCmd.Flags().StringVar(&logID, "id", "", "Identifiant unique du log (requis)")
	addLogCmd.Flags().StringVar(&logPath, "path", "", "Chemin vers le fichier de log (requis)")
	addLogCmd.Flags().StringVar(&logType, "type", "", "Type de log (requis)")
	addLogCmd.Flags().StringVar(&configFile, "file", "", "Fichier de configuration JSON (requis)")

	addLogCmd.MarkFlagRequired("id")
	addLogCmd.MarkFlagRequired("path")
	addLogCmd.MarkFlagRequired("type")
	addLogCmd.MarkFlagRequired("file")
}

func runAddLog() {
	newLog := config.LogConfig{
		ID:   logID,
		Path: logPath,
		Type: logType,
	}

	configs := []config.LogConfig{}
	
	if _, err := os.Stat(configFile); err == nil {
		existingConfigs, err := config.LoadConfig(configFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erreur lors du chargement de la configuration: %v\n", err)
			os.Exit(1)
		}
		
		for _, cfg := range existingConfigs {
			if cfg.ID == logID {
				fmt.Fprintf(os.Stderr, "Erreur: Un log avec l'ID '%s' existe déjà.\n", logID)
				os.Exit(1)
			}
		}
		
		configs = existingConfigs
	}
	
	configs = append(configs, newLog)
	
	if err := config.SaveConfig(configs, configFile); err != nil {
		fmt.Fprintf(os.Stderr, "Erreur lors de la sauvegarde de la configuration: %v\n", err)
		os.Exit(1)
	}
	
	fmt.Printf("Log '%s' ajouté avec succès à la configuration '%s'.\n", logID, configFile)
} 