FROM node:20-alpine

WORKDIR /app

COPY ./package.json ./
COPY ./yarn.lock ./
COPY *.ts ./
COPY *.json ./

RUN yarn

# COPY . ./

CMD [ "yarn", "dev" ]

EXPOSE 3000