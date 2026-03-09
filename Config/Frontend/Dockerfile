FROM node:22-alpine AS builder

WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
ARG BACK_URL=http://localhost:8080
ENV BACK_URL=$BACK_URL
RUN npm run generate

FROM node:22-alpine
WORKDIR /app
RUN npm install -g serve
COPY --from=builder /app/.output/public .

EXPOSE 3000
CMD ["serve", "-s", ".", "-l", "3000", "--single"]
