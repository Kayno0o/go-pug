FROM node:21-alpine as node

WORKDIR /app/client

COPY package*.json ./

RUN npm install -g pnpm
RUN pnpm install
RUN apk add --no-cache make

CMD ["make", "typescript"]
