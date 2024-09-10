FROM node:alpine

WORKDIR /Documents/admin
COPY ./admin .

RUN apk add --no-cache bash
RUN npm install -g http-server

RUN chmod +x ./wait_be.sh
CMD ["./wait_be.sh"]