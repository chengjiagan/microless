#! /bin/bash
set -ex

timestamp=$(date +%m%d%H%M)
mkdir ../data/$timestamp
ratefile="rate-test.csv"
ratio=1
service="socialnetwork"
api="mix"

load="../loader/main"

gateway="10.105.18.200:8080"
userid="../dataset/$service/user_ids.json"
movieid="../dataset/$service/movie_ids.json"
mode="file"
prewarn_duration=30
duration=180

output="../data/${timestamp}/load_ratefile.csv"
$load -service $service -api $api -addr $gateway -userid $userid -movieid $movieid -mode $mode -time $duration -file $ratefile -ratio $ratio -output $output
cp $ratefile "../data/${timestamp}/rate.csv"
