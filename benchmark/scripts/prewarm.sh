#! /bin/bash
set -ex

load="../loader/main"

service="socialnetwork"
gateway="10.101.6.50:8080"
userid="../dataset/$service/user_ids.json"
movieid="../dataset/$service/movie_ids.json"

$load -service $service -addr $gateway -userid $userid -movieid $movieid -mode "prewarm" -nthread 8
