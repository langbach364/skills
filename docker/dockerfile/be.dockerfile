FROM golang:latest

WORKDIR /Documents/backend
COPY ./backend .

WORKDIR /Documents/backend/
COPY ./backend/wait_mysql.sh .
RUN chmod +x wait_mysql.sh

COPY ./backend/start_sources.sh .
WORKDIR /Documents/backend/Connect_database
COPY ./backend/Connect_database/. .
RUN go mod tidy

WORKDIR /Documents/backend/System
COPY ./backend/System/. .
RUN go mod tidy

WORKDIR /Documents/backend
CMD ["./wait_mysql.sh"]
