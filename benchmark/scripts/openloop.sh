#! /bin/bash
set -ex

timestamp=$(date +%m%d%H%M)
mkdir ../data/$timestamp
qps=(50 80 100 130 150 180 200)
ratio=1
service="socialnetwork"

load="../loader/main"

gateway="gateway.social-network-kn.example.com"
userid="../dataset/$service/user_ids.json"
movieid="../dataset/$service/movie_ids.json"
mode="open"
prewarn_duration=60 # 1 mins
duration=120 # 2 mins

for n in ${qps[@]}; do
    # prewarm for boot up enough container
    $load -service $service -addr $gateway -userid $userid -movieid $movieid -mode $mode -time $prewarn_duration -rate $n -ratio $ratio -output /dev/null
    sleep 5
    # actual load test
    output="../data/${timestamp}/load_q${n}.csv"
    $load -service $service -addr $gateway -userid $userid -movieid $movieid -mode $mode -time $duration -rate $n -ratio $ratio -output $output
    sleep 60
done
