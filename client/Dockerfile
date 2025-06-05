FROM node:16-alpine as builder

WORKDIR /app

# Copier les fichiers de dépendances
COPY package.json package-lock.json ./

# Installer les dépendances
RUN npm ci

# Copier le code source
COPY . .

# Construire l'application
RUN npm run build

# Étape de production avec Nginx
FROM nginx:alpine

# Copier la configuration nginx personnalisée si nécessaire
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Copier les fichiers de build depuis l'étape précédente
COPY --from=builder /app/build /usr/share/nginx/html

# Exposer le port 80
EXPOSE 80

# Démarrer Nginx
CMD ["nginx", "-g", "daemon off;"]
