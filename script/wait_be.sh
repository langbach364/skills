#!/bin/bash

IP="172.21.0.4"
PORT="8080"
while ! ./wait-for-it.sh $IP:$PORT; do
    echo "Cổng $PORT từ địa chỉ $IP chưa mở..."
    sleep 15
done

echo "Cổng backend từ $IP:$PORT đã mở có thể chạy trang admin"
chmod +x ./start_web.sh
./start_web.sh