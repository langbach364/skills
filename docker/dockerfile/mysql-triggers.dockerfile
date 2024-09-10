FROM mysql:latest

ENV MYSQL_ROOT_PASSWORD="@ztegc4df9f4e"
WORKDIR /Documents/backend/Database

COPY ./backend/Database/triggers.sql /Database/triggers.sql
COPY ./backend/Database/add_data.sql /Database/add_data.sql
COPY ./backend/wait-for-it.sh .
COPY ./backend/Database/wait_mysql.sh .
COPY ./backend/Database/run_mysql.sh .

RUN chmod +x wait_mysql.sh
CMD ["./wait_mysql.sh"]