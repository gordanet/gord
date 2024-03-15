#!/bin/bash
rm -rf /tmp/gord-temp

kaspad --devnet --appdir=/tmp/gord-temp --profile=6061 --loglevel=debug &
GORD_PID=$!
gorD_KILLED=0
function killGordIfNotKilled() {
    if [ $GORD_KILLED -eq 0 ]; then
      kill $GORD_PID
    fi
}
trap "killgordIfNotKilled" EXIT

sleep 1

application-level-garbage --devnet -alocalhost:16611 -b blocks.dat --profile=7000
TEST_EXIT_CODE=$?

kill $GORD_PID

wait $GORD_PID
GORD_KILLED=1
GORD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Gord exit code: $GORD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $GORD_EXIT_CODE -eq 0 ]; then
  echo "application-level-garbage test: PASSED"
  exit 0
fi
echo "application-level-garbage test: FAILED"
exit 1
