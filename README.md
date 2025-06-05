# Social-Network

## Description

Ce projet est un réseau social de type Facebook-like développé dans le cadre pédagogique de Zone01. Il permet aux utilisateurs de créer des profils, publier des posts, commenter, aimer du contenu et suivre d'autres utilisateurs.

## Architecture technique

- **Frontend** : React.js
- **Backend** : Go (Golang) 
- **Base de données** : MySQL
- **Déploiement** : Docker (conteneurs séparés pour frontend, backend et base de données)

## Prérequis

- Docker et Docker Compose installés
- Git pour cloner le projet

## Installation et déploiement avec Docker

### 1. Cloner le dépôt

```bash
git clone <URL_DU_DÉPÔT>
cd social-network
```

### 2. Démarrer l'application avec Docker Compose

```bash
docker-compose up -d
```

Cette commande va:
- Construire les images Docker pour le frontend, backend et la base de données
- Créer et démarrer les conteneurs
- Configurer le réseau entre les services

### 3. Accéder à l'application

- **Frontend**: http://localhost
- **Backend API**: http://localhost:8080
- **Base de données MySQL**: localhost:3306 (accessible uniquement depuis les conteneurs ou localement)

## Structure des conteneurs Docker

### Frontend (React)
- Basé sur Node.js pour le build, puis Nginx pour servir l'application
- Optimisé pour la production
- Proxy des requêtes API vers le backend

### Backend (Go)
- Compilation et exécution du code Go dans un conteneur léger Alpine
- API RESTful exposée sur le port 8080
- Communication sécurisée avec la base de données

### Base de données (MySQL)
- Persistance des données via des volumes Docker
- Initialisation automatique du schéma de base de données
- Sécurisé avec identifiants configurables

## Tests

### Tests unitaires

Des tests unitaires ont été écrits pour le backend Go:

```bash
# Exécution des tests Go
cd server
go test ./...
```

### Test de déploiement Docker

Un script de test est fourni pour vérifier le déploiement Docker:

```bash
./test-docker.sh
```

Ce script vérifie:
- La construction des images Docker
- Le démarrage des services
- L'accessibilité du frontend et du backend
- La connexion à la base de données

## Fonctionnalités

- Authentification des utilisateurs (inscription, connexion, déconnexion)
- Création et gestion de profils utilisateurs
- Publication de posts avec texte et images
- Commentaires et réponses aux commentaires
- Système de likes pour les posts et commentaires
- Système de followers/following

## Sécurité

- Authentification par JWT (JSON Web Tokens)
- Mots de passe hashés (SHA-256)
- Cookies HTTP-only pour les sessions
- Middleware de vérification d'authentification

## Maintenance

### Arrêter les services

```bash
docker-compose down
```

### Voir les logs

```bash
docker-compose logs
# Ou pour un service spécifique
docker-compose logs backend
```

### Redémarrer un service

```bash
docker-compose restart backend
```

### Mise à jour

```bash
git pull
docker-compose down
docker-compose up -d --build
```
