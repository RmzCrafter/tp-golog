# LogAnalyzer - Outil d'Analyse de Logs Distribué

## 📋 Description

LogAnalyzer est un outil en ligne de commande (CLI) développé en Go qui permet d'analyser des fichiers de logs de manière distribuée et concurrente. Il utilise les goroutines pour traiter plusieurs logs en parallèle et offre une gestion robuste des erreurs.

## 🎯 Fonctionnalités

- **Analyse concurrente** : Traitement de plusieurs logs en parallèle via goroutines
- **Gestion d'erreurs robuste** : Erreurs personnalisées avec `errors.Is()` et `errors.As()`
- **Export JSON** : Sauvegarde des résultats au format JSON
- **Filtrage des résultats** : Possibilité de filtrer par statut (OK/FAILED)
- **Horodatage automatique** : Ajout de timestamps aux fichiers de sortie
- **Interface CLI intuitive** : Commandes et drapeaux clairs avec Cobra
- **Création automatique de répertoires** : Gestion intelligente des chemins de sortie

## 🏗️ Architecture

Le projet est organisé selon les bonnes pratiques Go :

```
loganalyzer/
├── main.go                 # Point d'entrée de l'application
├── cmd/                    # Commandes CLI (Cobra)
│   ├── root.go            # Commande racine
│   ├── analyze.go         # Commande d'analyse
│   └── addlog.go          # Commande d'ajout de log
├── internal/              # Packages internes
│   ├── config/            # Gestion des configurations JSON
│   │   └── config.go
│   ├── analyzer/          # Logique d'analyse et erreurs
│   │   ├── analyzer.go
│   │   └── errors.go
│   └── reporter/          # Export des résultats
│       └── reporter.go
├── example_config.json    # Fichier d'exemple de configuration
├── go.mod                 # Dépendances Go
└── README.md             # Documentation
```

## 🚀 Installation et Compilation

### Prérequis

- Go 1.19 ou supérieur
- Git

### Compilation

```bash

git clone <repository-url>
cd loganalyzer


go mod tidy


go build -o loganalyzer .


go install .
```








