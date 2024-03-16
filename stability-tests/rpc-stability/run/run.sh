#!/bin/bash
rm -rf /tmp/gord-temp

gord --devnet --appdir=/tmp/gord-temp --profile=6061 --loglevel=debug &
GORD_PID=$!

sleep 1

rpc-stability --devnet -p commands.json --profile=7000
TEST_EXIT_CODE=$?

kill $GORD_PID

wait $GORD_PID
GORD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Gord exit code: $GORD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $GORD_EXIT_CODE -eq 0 ]; then
  echo "rpc-stability test: PASSED"
  exit 0
fi
echo "rpc-stability test: FAILED"
exit 1
