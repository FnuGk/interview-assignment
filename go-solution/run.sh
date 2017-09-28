#!/usr/bin/env bash

rm -f test.db && cat ../sqldump.sql | sqlite3 test.db && rm -r out && go run main.go