# !/bin/bash
backend_file_start="./backend/start_sources.sh"
backend_file_wait_mysql="./backend/wait_mysql.sh"
backend_file_support="./backend/wait-for-it.sh"

frontend_file_start="./frontend/start_web.sh"
frontend_file_support="./frontend/wait-for-it.sh"
frontend_file_wait_be="./frontend/wait_be.sh"

traefik_acme="./docker/dockerfile/traefik/acme.json"

admin_file_start="./admin/start_web.sh"
admin_file_support="./admin/wait-for-it.sh"
admin_file_wait_be="./admin/wait_be.sh"

chmod 755 $backend_file_start
chmod 755 $backend_file_wait_mysql
chmod 755 $backend_file_support
chmod 755 $frontend_file_start
chmod 755 $frontend_file_support
chmod 755 $frontend_file_wait_be
chmod 600 $traefik_acme
chmod 755 $admin_file_start
chmod 755 $admin_file_support
chmod 755 $admin_file_wait_be