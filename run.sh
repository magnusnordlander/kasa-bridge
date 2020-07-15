#!/bin/bash

echo "Starting with a 5 second nap"
sleep 5

envsubst < /app/nginx.conf > /etc/nginx/conf.d/default.conf
nginx

until /app/main | ffmpeg -r 15 -i - -c copy -f flv $STREAM_URL; do
	echo "Camera crashed with $?.  Respawning.." >&2
    sleep 10
done