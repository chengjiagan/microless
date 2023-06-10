redis-cli -h 139.224.190.177 -n 7 SET post-storage true
redis-cli -h 139.224.190.177 -n 7 SET home-timeline true
redis-cli -h 139.224.190.177 -n 7 SET user-timeline true
redis-cli -h 139.224.190.177 -n 7 PUBLISH post-storage true
redis-cli -h 139.224.190.177 -n 7 PUBLISH home-timeline true
redis-cli -h 139.224.190.177 -n 7 PUBLISH user-timeline true