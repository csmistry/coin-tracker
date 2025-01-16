FROM node:22 AS build

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY src/ ./src/
COPY public/ ./public/

RUN npm run build

FROM nginx:alpine 

COPY --from=build /app/build /usr/share/nginx/html

EXPOSE 80

# Run Nginx to serve the React app
CMD ["nginx", "-g", "daemon off;"]