# Architecture MVC du Backend

```
server/
u251cu2500u2500 controllers/            # Contrôleurs (C de MVC) - Gestion des requêtes et réponses
u2502   u251cu2500u2500 auth.go            # Gestion de l'authentification (login, signup, logout)
u2502   u251cu2500u2500 comOfComs.go        # Gestion des réponses aux commentaires
u2502   u251cu2500u2500 comments.go         # Gestion des commentaires sur les posts
u2502   u251cu2500u2500 likes.go            # Gestion des likes sur les posts
u2502   u251cu2500u2500 likesComs.go        # Gestion des likes sur les commentaires
u2502   u251cu2500u2500 posts.go            # Gestion des publications
u2502   u2514u2500u2500 upload.go           # Gestion des téléchargements de fichiers
u2502
u251cu2500u2500 models/                # Modèles (M de MVC) - Structure des données et accès à la base de données
u2502   u251cu2500u2500 comOfComs.go        # Modèle des réponses aux commentaires
u2502   u251cu2500u2500 comments.go         # Modèle des commentaires
u2502   u251cu2500u2500 db.go               # Configuration de la connexion à la base de données
u2502   u251cu2500u2500 likes.go            # Modèle des likes
u2502   u251cu2500u2500 likesComs.go        # Modèle des likes de commentaires
u2502   u251cu2500u2500 posts.go            # Modèle des publications
u2502   u251cu2500u2500 relationships.go    # Modèle des relations entre utilisateurs (followers)
u2502   u2514u2500u2500 users.go            # Modèle des utilisateurs
u2502
u251cu2500u2500 routes/                # Routage des requêtes HTTP
u2502   u251cu2500u2500 auth.go             # Routes d'authentification
u2502   u251cu2500u2500 comOfComs.go         # Routes des réponses aux commentaires
u2502   u251cu2500u2500 comments.go          # Routes des commentaires
u2502   u251cu2500u2500 likes.go             # Routes des likes
u2502   u251cu2500u2500 likesComs.go         # Routes des likes de commentaires
u2502   u2514u2500u2500 posts.go             # Routes des publications
u2502
u251cu2500u2500 middleware/            # Middleware pour l'authentification et la sécurité
u2502   u251cu2500u2500 checker.go          # Vérification de l'authentification
u2502   u2514u2500u2500 logging.go          # Journalisation des requêtes
u2502
u251cu2500u2500 db/                    # Scripts et configuration de la base de données
u2502   u2514u2500u2500 database.sql        # Schéma de la base de données MySQL
u2502
u251cu2500u2500 main.go                # Point d'entrée de l'application
u251cu2500u2500 go.mod                 # Gestion des dépendances
u2514u2500u2500 go.sum                 # Checksums des dépendances
```

## Description de l'architecture MVC

### Modèle (M)

Le répertoire `models/` contient la logique d'accès aux données et les structures qui représentent les entités de l'application :

- Définition des structures de données (users, posts, comments, etc.)
- Fonctions pour interagir avec la base de données MySQL
- Validation des données

### Vue (V)

Dans cette architecture backend API, il n'y a pas de "Vue" traditionnelle comme dans une application MVC complète. Le frontend React est responsable de la partie Vue. Le backend renvoie des données au format JSON qui seront interprétées et affichées par le frontend.

### Contrôleur (C)

Le répertoire `controllers/` contient la logique métier qui traite les requêtes HTTP :

- Récupération des données des requêtes
- Appel aux fonctions du modèle pour accéder aux données
- Transformation des données si nécessaire
- Génération des réponses HTTP appropriées

### Routage

Le répertoire `routes/` définit les endpoints de l'API et associe chaque URL à un contrôleur spécifique.

### Middleware

Le répertoire `middleware/` contient des fonctions intermédiaires qui s'exécutent entre la requête et le contrôleur :

- Authentification et autorisation
- Journalisation
- Gestion des erreurs

## Flux d'une requête

1. Une requête HTTP arrive au serveur (main.go)
2. Le routeur (routes/) identifie le contrôleur approprié
3. Les middlewares (middleware/) vérifient l'authentification si nécessaire
4. Le contrôleur (controllers/) traite la requête
5. Le contrôleur utilise les modèles (models/) pour accéder aux données
6. Le contrôleur génère une réponse JSON
7. La réponse est renvoyée au client (frontend React)
