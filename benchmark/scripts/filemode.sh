#! /bin/bash
set -ex

timestamp=$(date +%m%d%H%M)
mkdir ../data/$timestamp
ratefile="rate.csv"
ratio=1
service="pingpong"

load="../loader/main"

gateway="10.111.64.109:8080"
userid="../dataset/$service/user_ids.json"
movieid="../dataset/$service/movie_ids.json"
mode="file"
prewarn_duration=30
duration=180

output="../data/${timestamp}/load_ratefile.csv"
$load -service $service -addr $gateway -userid $userid -movieid $movieid -mode $mode -time $duration -file $ratefile -ratio $ratio -output $output
cp $ratefile "../data/${timestamp}/rate.csv"
