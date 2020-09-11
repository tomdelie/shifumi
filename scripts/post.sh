#!bin/sh

ab -n 100 -c 100 -v 2 -T application/json -p post_data.json http://localhost:8081/players/1/move
