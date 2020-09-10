#!bin/sh

curl -X POST -d "{\"move\": \"paper\"}" http://localhost:8081/players/1/move
echo "\n"