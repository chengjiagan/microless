#! /bin/bash
set -ex

timestamp=$(date +%m%d%H%M)
mkdir ../data/$timestamp
qps=(200)
ratio=1
service="socialnetwork"
api="mix"

load="../loader/main"

gateway="10.98.64.87:8080"
userid="../dataset/$service/user_ids.json"
movieid="../dataset/$service/movie_ids.json"
mode="open"
prewarn_duration=60 # 1 mins
duration=120 # 2 mins

for n in ${qps[@]}; do
    # prewarm for boot up enough container
    # $load -service $service -api $api -addr $gateway -userid $userid -movieid $movieid -mode $mode -time $prewarn_duration -rate $n -ratio $ratio -output /dev/null
    # sleep 5
    # actual load test
    output="../data/${timestamp}/load_q${n}.csv"
    $load -service $service -api $api -addr $gateway -userid $userid -movieid $movieid -mode $mode -time $duration -rate $n -ratio $ratio -output $output
    sleep 60
done
