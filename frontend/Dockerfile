# Build stage
FROM node:lts-alpine3.18 as builder
WORKDIR /app
COPY . .
RUN yarn install
RUN yarn build

# Release stage
FROM nginx:alpine3.17-slim
WORKDIR /app
COPY --from=builder /app/dist /usr/share/nginx/html

# Start NGINX when the container starts
CMD ["nginx", "-g", "daemon off;"]
