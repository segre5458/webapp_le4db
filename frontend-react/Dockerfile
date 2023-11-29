FROM node:alpine

RUN mkdir -p /webapp_le4db
COPY . /webapp_le4db
WORKDIR /webapp_le4db/frontend-react

RUN npm install

RUN npm run build

EXPOSE 3000

CMD ["npm", "start"]
