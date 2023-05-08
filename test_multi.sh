#!/bin/zsh
for i in $(seq 1 3333); do
  req_n=$((i / 100))
  curl -s -o /dev/null "http://localhost:1233/fib?n=${req_n}";
done &
for i in $(seq 3334 6666); do
  req_n=$((i / 100))
  curl -s -o /dev/null "http://localhost:1234/fib?n=${req_n}";
done &
for i in $(seq 6667 10000); do
  req_n=$((i / 100))
  curl -s -o /dev/null "http://localhost:1235/fib?n=${req_n}";
done