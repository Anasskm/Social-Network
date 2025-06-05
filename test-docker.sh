#!/bin/bash

# Script de test pour le déploiement Docker du réseau social

echo "=== Test de déploiement Docker pour le réseau social ==="

# Vérifier que Docker et Docker Compose sont installés
if ! command -v docker &> /dev/null; then
    echo "Erreur: Docker n'est pas installé"
    exit 1
fi

if ! command -v docker-compose &> /dev/null; then
    echo "Erreur: Docker Compose n'est pas installé"
    exit 1
fi

echo "✓ Docker et Docker Compose sont installés"

# Construire les images Docker
echo "Construction des images Docker..."
docker-compose build

if [ $? -ne 0 ]; then
    echo "Erreur: La construction des images Docker a échoué"
    exit 1
fi

echo "✓ Images Docker construites avec succès"

# Démarrer les services
echo "Démarrage des services Docker..."
docker-compose up -d

if [ $? -ne 0 ]; then
    echo "Erreur: Le démarrage des services Docker a échoué"
    exit 1
fi

echo "✓ Services Docker démarrés avec succès"

# Attendre que les services soient prêts
echo "Attente du démarrage complet des services..."
sleep 15

# Tester le service backend
echo "Test du service backend..."
curl -s -o /dev/null -w "%{http_code}" http://localhost:8080

if [ $? -ne 0 ]; then
    echo "Erreur: Le service backend n'est pas accessible"
    docker-compose logs backend
    docker-compose down
    exit 1
fi

echo "✓ Service backend accessible"

# Tester le service frontend
echo "Test du service frontend..."
curl -s -o /dev/null -w "%{http_code}" http://localhost

if [ $? -ne 0 ]; then
    echo "Erreur: Le service frontend n'est pas accessible"
    docker-compose logs frontend
    docker-compose down
    exit 1
fi

echo "✓ Service frontend accessible"

# Tester la connexion à la base de données depuis le backend
echo "Test de la connexion à la base de données..."
docker-compose exec backend wget -qO- http://localhost:8080/health

if [ $? -ne 0 ]; then
    echo "Erreur: La connexion à la base de données depuis le backend a échoué"
    docker-compose logs database
    docker-compose down
    exit 1
fi

echo "✓ Connexion à la base de données réussie"

echo "=== Tous les tests ont réussi ! ==="
echo "Les services sont en cours d'exécution et accessibles :"
echo "- Frontend: http://localhost"
echo "- Backend: http://localhost:8080"
echo "- Base de données MySQL: localhost:3306"

echo "Pour arrêter les services, exécutez : docker-compose down"
