wget "https://github.com/docker/compose/releases/download/v2.27.0/docker-compose-linux-x86_64"

sudo -s
mv docker-compose-linux-x86_64 docker-compose
mv docker-compose /usr/bin/
chmod +x /usr/bin/docker-compose