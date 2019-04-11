#!/bin/bash
set -e

# start x server
X -config ./scripts/assets/dummydisplay.conf &
#set display to use
export DISPLAY=:0

statusfile=$(mktemp)
logfile=$(mktemp)
echo "Starting texts in Xterm"
xterm -e sh -c 'go test -v -count=1 ./... > '$logfile'; echo $? > '$statusfile
echo "Tests finished"
status=$(cat $statusfile)
rm $statusfile

echo "Tests returned exit code: $status"

if [[ $status == "0" ]]; then
  cat $logfile
  echo "Tests passed"
else
  cat $logfile
  echo "Tests FAILED"
  go version
  exit 1
fi 
