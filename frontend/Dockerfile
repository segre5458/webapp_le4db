FROM node:16.19-alpine

RUN mkdir -p /app
WORKDIR /app

ADD package.json package-locak.json ./
RUN npm ci

ADD . ./

CMD ["npm", "run", "dev"]
