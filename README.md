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

## 📝 Configuration

Format JSON avec tableau de logs :

```json
[
  {
    "id": "web-server",
    "path": "/var/log/nginx/access.log",
    "type": "nginx"
  },
  {
    "id": "app-backend",
    "path": "/var/log/app/application.log",
    "type": "application"
  }
]
```

**Champs requis :** `id` (identifiant unique), `path` (chemin fichier), `type` (type de log)

## 🔧 Utilisation

### Commande analyze

```bash
# Analyse basique
./loganalyzer analyze --config example_config.json

# Avec export et filtrage
./loganalyzer analyze -c config.json -o results.json --status OK
./loganalyzer analyze -c config.json -o results.json --status FAILED

# Avec horodatage automatique
./loganalyzer analyze -c config.json -o results.json --timestamp
```

### Commande add-log

```bash
# Ajouter un log
./loganalyzer add-log --id web-logs --path /var/log/nginx.log --type nginx --file config.json

# Créer nouvelle configuration
./loganalyzer add-log --id app-log --path /tmp/app.log --type application --file new_config.json
```

**Options :** `--config/-c` (fichier config), `--output/-o` (sortie JSON), `--status` (filtre OK/FAILED), `--timestamp` (horodatage)

## 📊 Format de sortie

```json
[
  {
    "id": "web-server",
    "status": "OK"
  },
  {
    "id": "app-backend",
    "status": "FAILED",
    "error": "Erreur lors de l'analyse du log app-backend"
  }
]
```

**Statuts :** `OK` (succès), `FAILED` (erreur avec détails)

## 🐛 Dépannage

**Fichier config introuvable :**

```bash
ls -la config.json  # Vérifier existence
./loganalyzer analyze --config /chemin/absolu/config.json
```

**Statut invalide :** Utiliser `--status OK` ou `--status FAILED` (respecter la casse)

**ID dupliqué :**

```bash
cat config.json | grep -o '"id": "[^"]*"'  # Lister IDs existants
```

## 🔧 Développement

```bash
go fmt ./...                              # Formatage
go vet ./...                              # Vérification
go test ./...                             # Tests
go build -ldflags "-s -w" -o loganalyzer . # Build optimisé
```

**Extensions :** Nouveaux types d'erreurs (`internal/analyzer/errors.go`), nouvelles commandes (`cmd/`), formats d'export (`internal/reporter/`)

## 📈 Performance

- **Concurrence** : Goroutines pour traitement parallèle efficace
- **Mémoire** : Collecte résultats via channels sans stockage excessif
- **I/O** : Création répertoires à la demande
- **Démarrage** : < 50ms, analyse proportionnelle au fichier le plus lent

## 📄 Licence

Projet TP académique - Go et programmation concurrente.
