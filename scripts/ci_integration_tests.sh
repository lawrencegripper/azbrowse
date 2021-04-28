#!/usr/bin/env bash
set -e

cd `dirname $0`

Xvfb :99 -ac -screen 0 "$XVFB_RES" -nolisten tcp $XVFB_ARGS &
XVFB_PROC=$!
sleep 1
export DISPLAY=:99
exitcodefile=$(mktemp)
logfile=$(mktemp)

echo "Run make integration in Xterm"
xterm -e sh -c 'make integration > '"$logfile"'; echo $? > '"$exitcodefile"
echo "Tests finished"
cat $exitcodefile
exitcode=$(cat "$exitcodefile")
rm "$exitcodefile"

if [[ $exitcode == "0" ]]; then
  cat "$logfile"

  echo "Tests passed"
else
  cat "$logfile"
  go version
  
  echo "Tests returned exit code: $exitcode"
  echo "Tests FAILED. Logs:"
  exit 1
fi 
kill $XVFB_PROC