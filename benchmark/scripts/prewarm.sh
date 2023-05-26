#! /bin/bash
set -ex

src="../loader/main.go"
load="./main"
go build -o $load $src

service="socialnetwork"
gateway="10.101.6.50:8080"
userid="../dataset/$service/user_ids.json"
movieid="../dataset/$service/movie_ids.json"

$load -service $service -addr $gateway -userid $userid -movieid $movieid -mode "prewarm" -nthread 8
rm $load
