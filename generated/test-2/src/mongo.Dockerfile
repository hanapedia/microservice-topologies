FROM --platform=linux/amd64 mongo:4

ENV MONGO_INITDB_ROOT_USERNAME=root
ENV MONGO_INITDB_ROOT_PASSWORD=example
COPY ./mongoScript.js /docker-entrypoint-initdb.d/mongoScript.js

EXPOSE 27017
