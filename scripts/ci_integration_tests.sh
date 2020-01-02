#!/bin/bash
set -e

echo ""
echo "-------> INTEGRATION TESTS"
echo ""

echo "Starting Xvfb"

Xvfb :99 -ac -screen 0 "$XVFB_RES" -nolisten tcp $XVFB_ARGS &
XVFB_PROC=$!
sleep 1
export DISPLAY=:99

statusfile=$(mktemp)
logfile=$(mktemp)

echo "Starting tests in Xterm"
xterm -e sh -c 'go test -v -count=1 ./... > '"$logfile"'; echo $? > '"$statusfile"
echo "Tests finished"
status=$(cat "$statusfile")
rm "$statusfile"

echo "Tests returned exit code: $status"

if [[ $status == "0" ]]; then
  cat "$logfile"
  echo "Tests passed"
else
  cat "$logfile"
  echo "Tests FAILED"
  go version
  exit 1
fi 

kill $XVFB_PROC

echo "S"
