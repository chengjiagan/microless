#! /bin/bash
set -ex

timestamp=$(date +%m%d%H%M)
mkdir ../data/$timestamp
rthread=(0 1 2 3 4)
wthread=(0)
service="socialnetwork"
api="mix"
load="../loader/main"

gateway="127.0.0.1:80"
userid="../dataset/$service/user_ids.json"
movieid="../dataset/$service/movie_ids.json"
mode="close"
prewarn_duration=60 # 1 mins
duration=120 # 2 mins

for r in ${rthread[@]}; do
    for w in ${wthread[@]}; do
        # prewarm for boot up enough container
        $load -service $service -api $api -addr $gateway -userid $userid -movieid $movieid -mode $mode -time $prewarn_duration -rthread $r -wthread $w -output /dev/null
        sleep 5
        # actual load test
        output="../data/${timestamp}/load_r${r}w${w}.csv"
        $load -service $service -api $api -addr $gateway -userid $userid -movieid $movieid -mode $mode -time $duration -rthread $r -wthread $w -output $output
        sleep 60
    done
done
