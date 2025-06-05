FROM golang:1.18-alpine AS builder

WORKDIR /app

# Copier les fichiers de dépendances
COPY go.mod go.sum ./
RUN go mod download

# Copier le code source
COPY . .

# Compiler l'application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Étape finale avec une image minimaliste
FROM alpine:latest

WORKDIR /app

# Installer les certificats pour les requêtes HTTPS
RUN apk --no-cache add ca-certificates

# Copier le binaire compilé
COPY --from=builder /app/main /app/

# Exposer le port utilisé par l'application
EXPOSE 8080

# Lancer l'application
CMD ["/app/main"]
