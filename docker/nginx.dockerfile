FROM node:16.15-alpine3.15 as build

WORKDIR app/

COPY ./web/package.json .
COPY ./web/package-lock.json .

RUN npm install && mv node_modules ../

COPY ./web .

RUN npm run build

FROM nginx:1.22-alpine

COPY --from=build /app/dist /app/
COPY ./nginx.conf /etc/nginx/nginx.conf

RUN mkdir /etc/nginx/logs
RUN touch /etc/nginx/logs/error.log

