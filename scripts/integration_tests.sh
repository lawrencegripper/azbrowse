#!/bin/bash
set -e

statusfile=$(mktemp)
logfile=$(mktemp)
xterm -e sh -c 'make test > '$logfile'; echo $? > '$statusfile
status=$(cat $statusfile)
rm $statusfile

echo "Tests returned exit code: $status"

if [[ $status == "0" ]]; then
  cat $logfile
  echo "Tests passed"
else
  cat $logfile
  echo "Tests FAILED"
  exit 1
fi 
