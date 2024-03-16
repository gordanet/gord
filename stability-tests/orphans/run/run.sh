#!/bin/bash
rm -rf /tmp/gord-temp

gord --simnet --appdir=/tmp/gordd-temp --profile=6061 &
GORD_PID=$!

sleep 1

orphans --simnet -alocalhost:16511 -n20 --profile=7000
TEST_EXIT_CODE=$?

kill $GORD_PID

wait $GORD_PID
KASPAD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Gord exit code: $GORD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $GORD_EXIT_CODE -eq 0 ]; then
  echo "orphans test: PASSED"
  exit 0
fi
echo "orphans test: FAILED"
exit 1
