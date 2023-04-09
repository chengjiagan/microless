#! /bin/bash
set -ex

src="../load_social_network/main.go"
load="./main"
go build -o $load $src

gateway="gateway.social-network-kn.example.com"
userid="../dataset/socialnetwork/user_ids.json"

$load -addr $gateway -userid $userid -mode "prewarm" -nthread 8

rm $load
