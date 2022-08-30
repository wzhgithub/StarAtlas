for ((i=1;i<=1000;i++)); do curl -i -X POST -H "Content-Type: application/json; charset=utf-8" -d@failure_vmc.json \
"http://127.0.0.1:9999/vmc/do_failure_over"; sleep 1s; done