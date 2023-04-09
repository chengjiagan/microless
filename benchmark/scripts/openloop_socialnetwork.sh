#! /bin/bash
set -ex

timestamp=$(date +%m%d%H%M)
mkdir ../data/$timestamp
qps=(30 60 90 120 150 180)
ratio=1

src="../load_social_network/main.go"
load="./main"
go build -o $load $src

gateway="gateway.social-network-kn.example.com"
userid="../dataset/socialnetwork/user_ids.json"
mode="open"
prewarn_duration=60 # 1 mins
duration=120 # 2 mins

for n in ${qps[@]}; do
    # prewarm for boot up enough container
    $load -addr $gateway -userid $userid -mode $mode -time $prewarn_duration -rate $n -ratio $ratio -output /dev/null
    sleep 5
    # actual load test
    output="../data/${timestamp}/load_q${n}.csv"
    $load -addr $gateway -userid $userid -mode $mode -time $duration -rate $n -ratio $ratio -output $output
    sleep 60
done
rm $load
