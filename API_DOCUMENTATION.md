# Documentation de l'API du ru00e9seau social

Cette documentation du00e9crit les endpoints disponibles dans l'API backend du ru00e9seau social, avec des exemples d'utilisation via cURL.

## Table des matiu00e8res

1. [Authentification](#authentification)
2. [Posts](#posts)
3. [Commentaires](#commentaires)
4. [Ru00e9ponses aux commentaires](#ru00e9ponses-aux-commentaires)
5. [Likes](#likes)
6. [Relations (followers)](#relations)

## Base URL

Toutes les URLs sont relatives u00e0 : `http://localhost:8080`

## Authentification

### Inscription

```
POST /signup
```

**Corps de la requu00eate :**

```json
{
  "username": "utilisateur",
  "email": "utilisateur@example.com",
  "password": "motdepasse123",
  "firstname": "Pru00e9nom",
  "lastname": "Nom",
  "dateOfBirth": "1990-01-01",
  "city": "Paris"
}
```

**Exemple avec cURL :**

```bash
curl -X POST http://localhost:8080/signup \
  -H "Content-Type: application/json" \
  -d '{
    "username": "utilisateur",
    "email": "utilisateur@example.com",
    "password": "motdepasse123",
    "firstname": "Pru00e9nom",
    "lastname": "Nom",
    "dateOfBirth": "1990-01-01",
    "city": "Paris"
  }'
```

### Connexion

```
POST /login
```

**Corps de la requu00eate :**

```json
{
  "username": "utilisateur",
  "password": "motdepasse123"
}
```

**Exemple avec cURL :**

```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -c cookies.txt \
  -d '{
    "username": "utilisateur",
    "password": "motdepasse123"
  }'
```

> Note: Le `-c cookies.txt` sauvegarde le cookie de session dans un fichier pour ru00e9utilisation.

### Du00e9connexion

```
GET /logout
```

**Exemple avec cURL :**

```bash
curl -X GET http://localhost:8080/logout \
  -b cookies.txt
```

> Note: Le `-b cookies.txt` utilise le cookie de session sauvegardu00e9 pru00e9cu00e9demment.

## Posts

### Cru00e9er un post

```
POST /api/posts
```

**Corps de la requu00eate :**

```json
{
  "desc": "Contenu du post",
  "img": "url_de_l_image.jpg"
}
```

**Exemple avec cURL :**

```bash
curl -X POST http://localhost:8080/api/posts \
  -H "Content-Type: application/json" \
  -b cookies.txt \
  -d '{
    "desc": "Contenu du post",
    "img": "url_de_l_image.jpg"
  }'
```

### Ru00e9cupu00e9rer tous les posts

```
GET /api/posts
```

**Exemple avec cURL :**

```bash
curl -X GET http://localhost:8080/api/posts \
  -b cookies.txt
```

### Supprimer un post

```
DELETE /api/posts/{postId}
```

**Exemple avec cURL :**

```bash
curl -X DELETE http://localhost:8080/api/posts/1 \
  -b cookies.txt
```

## Commentaires

### Ajouter un commentaire u00e0 un post

```
POST /api/comments
```

**Corps de la requu00eate :**

```json
{
  "desc": "Contenu du commentaire",
  "postId": 1
}
```

**Exemple avec cURL :**

```bash
curl -X POST http://localhost:8080/api/comments \
  -H "Content-Type: application/json" \
  -b cookies.txt \
  -d '{
    "desc": "Contenu du commentaire",
    "postId": 1
  }'
```

### Ru00e9cupu00e9rer les commentaires d'un post

```
GET /api/comments/{postId}
```

**Exemple avec cURL :**

```bash
curl -X GET http://localhost:8080/api/comments/1 \
  -b cookies.txt
```

## Ru00e9ponses aux commentaires

### Ajouter une ru00e9ponse u00e0 un commentaire

```
POST /api/comofcoms
```

**Corps de la requu00eate :**

```json
{
  "desc": "Ru00e9ponse au commentaire",
  "commentId": 1
}
```

**Exemple avec cURL :**

```bash
curl -X POST http://localhost:8080/api/comofcoms \
  -H "Content-Type: application/json" \
  -b cookies.txt \
  -d '{
    "desc": "Ru00e9ponse au commentaire",
    "commentId": 1
  }'
```

### Ru00e9cupu00e9rer les ru00e9ponses d'un commentaire

```
GET /api/comofcoms/{commentId}
```

**Exemple avec cURL :**

```bash
curl -X GET http://localhost:8080/api/comofcoms/1 \
  -b cookies.txt
```

## Likes

### Aimer un post

```
POST /api/likes
```

**Corps de la requu00eate :**

```json
{
  "postId": 1
}
```

**Exemple avec cURL :**

```bash
curl -X POST http://localhost:8080/api/likes \
  -H "Content-Type: application/json" \
  -b cookies.txt \
  -d '{
    "postId": 1
  }'
```

### Supprimer un like d'un post

```
DELETE /api/likes/{postId}
```

**Exemple avec cURL :**

```bash
curl -X DELETE http://localhost:8080/api/likes/1 \
  -b cookies.txt
```

### Aimer un commentaire

```
POST /api/likescoms
```

**Corps de la requu00eate :**

```json
{
  "commentId": 1
}
```

**Exemple avec cURL :**

```bash
curl -X POST http://localhost:8080/api/likescoms \
  -H "Content-Type: application/json" \
  -b cookies.txt \
  -d '{
    "commentId": 1
  }'
```

## Relations

### Suivre un utilisateur

```
POST /api/relationships
```

**Corps de la requu00eate :**

```json
{
  "followedUserId": 2
}
```

**Exemple avec cURL :**

```bash
curl -X POST http://localhost:8080/api/relationships \
  -H "Content-Type: application/json" \
  -b cookies.txt \
  -d '{
    "followedUserId": 2
  }'
```

### Ne plus suivre un utilisateur

```
DELETE /api/relationships/{userId}
```

**Exemple avec cURL :**

```bash
curl -X DELETE http://localhost:8080/api/relationships/2 \
  -b cookies.txt
```

### Ru00e9cupu00e9rer les followers d'un utilisateur

```
GET /api/relationships/followers/{userId}
```

**Exemple avec cURL :**

```bash
curl -X GET http://localhost:8080/api/relationships/followers/1 \
  -b cookies.txt
```

### Ru00e9cupu00e9rer les utilisateurs suivis par un utilisateur

```
GET /api/relationships/following/{userId}
```

**Exemple avec cURL :**

```bash
curl -X GET http://localhost:8080/api/relationships/following/1 \
  -b cookies.txt
```

## Notes de su00e9curitu00e9

- Toutes les routes (sauf /signup et /login) nu00e9cessitent une authentification valide
- Les requu00eates sont limitu00e9es en taille et en nombre pour u00e9viter les attaques DoS
- Les cookies sont configuru00e9s comme HTTP-only pour u00e9viter les attaques XSS
- Les tokens JWT expirent apru00e8s 24 heures
