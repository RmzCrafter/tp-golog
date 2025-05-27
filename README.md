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

## 📖 Utilisation

### Commande `analyze`

Analyse les fichiers de logs spécifiés dans un fichier de configuration JSON.

```bash

./loganalyzer analyze --config example_config.json


./loganalyzer analyze -c example_config.json -o results.json


./loganalyzer analyze -c example_config.json --status FAILED


./loganalyzer analyze -c example_config.json -o results.json --timestamp
```

#### Drapeaux disponibles

- `-c, --config` : Chemin vers le fichier de configuration JSON (requis)
- `-o, --output` : Chemin vers le fichier de sortie JSON (optionnel)
- `--status` : Filtrer par statut (OK ou FAILED)
- `--timestamp` : Ajouter un horodatage au fichier de sortie

### Commande `add-log`

Ajoute une nouvelle configuration de log au fichier de configuration existant.

```bash

./loganalyzer add-log --id web-server-1 --path /var/log/nginx.log --type nginx --file config.json

# Ajouter un log d'application
./loganalyzer add-log --id app-backend --path /var/log/app.log --type custom-app --file config.json
```

#### Drapeaux disponibles

- `--id` : Identifiant unique pour le log (requis)
- `--path` : Chemin vers le fichier de log (requis)
- `--type` : Type du log (requis)
- `--file` : Chemin vers le fichier de configuration (défaut: config.json)

## 📄 Format de Configuration

Le fichier de configuration doit être au format JSON et contenir un tableau d'objets log :

```json
[
  {
    "id": "web-server-1",
    "path": "/var/log/nginx/access.log",
    "type": "nginx-access"
  },
  {
    "id": "app-backend-2",
    "path": "/var/log/my_app/errors.log",
    "type": "custom-app"
  }
]
```

### Champs requis

- `id` : Identifiant unique pour le log
- `path` : Chemin (absolu ou relatif) vers le fichier de log
- `type` : Type de log (libre, pour classification)

## 📊 Format de Sortie

Les résultats sont exportés au format JSON :

```json
[
  {
    "log_id": "web-server-1",
    "file_path": "/var/log/nginx/access.log",
    "status": "OK",
    "message": "Analyse terminée avec succès.",
    "error_details": ""
  },
  {
    "log_id": "invalid-path",
    "file_path": "/non/existent/log.log",
    "status": "FAILED",
    "message": "Fichier introuvable.",
    "error_details": "fichier introuvable: /non/existent/log.log"
  }
]
```

## 🔧 Gestion des Erreurs

L'application implémente deux types d'erreurs personnalisées :

1. **FileNotFoundError** : Fichier introuvable ou inaccessible
2. **ParsingError** : Erreur lors du parsing du fichier (simulée à 10%)

Ces erreurs sont gérées avec `errors.Is()` et `errors.As()` pour une identification précise du type d'erreur.

## ⚡ Concurrence

L'application utilise :

- **Goroutines** : Une goroutine par fichier de log à analyser
- **WaitGroup** : Synchronisation des goroutines
- **Channels** : Collecte sécurisée des résultats
- **Mutex** : Protection des données partagées

## 🎁 Fonctionnalités Bonus Implémentées

1. **Gestion des dossiers d'exportation** : Création automatique des répertoires manquants
2. **Horodatage des exports JSON** : Format AAMMJJ (ex: 241124_results.json)
3. **Commande add-log** : Ajout manuel de configurations
4. **Filtrage des résultats** : Par statut OK/FAILED

## 🧪 Tests et Exemples

### Test rapide

```bash

go build -o loganalyzer .


./loganalyzer analyze -c example_config.json


./loganalyzer analyze -c example_config.json -o test_results.json


./loganalyzer add-log --id test-log --path ./README.md --type documentation --file example_config.json
```

## 👥 Équipe de Développement

- **Développeur Principal** : [Votre Nom]
- **Technologies utilisées** : Go, Cobra CLI framework
- **Version** : 1.0.0

## 📝 Documentation du Code

Le code est entièrement documenté avec :

- Commentaires pour toutes les fonctions publiques
- Documentation des structures et interfaces
- Exemples d'utilisation dans les commentaires
- Messages d'erreur explicites

## 🔍 Algorithme de Simulation

L'analyse simule un traitement réaliste :

1. Vérification de l'existence du fichier
2. Délai aléatoire de traitement (50-200ms)
3. Simulation d'erreur de parsing (10% de chance)
4. Génération du rapport de résultat

## 🐛 Gestion des Erreurs

L'application gère proprement :

- Fichiers de configuration invalides
- Fichiers de logs introuvables
- Permissions insuffisantes
- Erreurs de création de répertoires
- Conflits d'ID dans la configuration

## 📈 Performance

- Traitement concurrent pour optimiser le temps d'exécution
- Gestion mémoire efficace avec channels bufferisés
- Pas de blocage grâce aux WaitGroups
- Messages de progression en temps réel

---

**Note** : Ce projet a été développé dans le cadre d'un TP académique pour démontrer la maîtrise des concepts de concurrence, gestion d'erreurs et développement CLI en Go.
