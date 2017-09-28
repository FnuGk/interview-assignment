#!/usr/bin/env bash

set -e # Stops the script on first failure

dbName="test.db"

rm -f "$dbName" && cat ../sqldump.sql | sqlite3 "$dbName" && \
sqlite3 -csv "$dbName" "SELECT id,firstName,lastName,email FROM users" > users.csv && \
nUsers="$(cat users.csv | wc -l | cut -d' ' -f 6)" && \
sqlite3 "$dbName" "DELETE FROM users WHERE $nUsers=(SELECT count(*) FROM users)" && \
echo "SUCCESS"
