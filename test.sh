#!/bin/zsh
for i in $(seq 1 10000); do
  req_n=$((i / 100))
  curl -s -o /dev/null "http://localhost:1233/fib?n=${req_n}";
done