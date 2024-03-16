#!/bin/bash
rm -rf /tmp/gord-temp

NUM_CLIENTS=128
gord --devnet --appdir=/tmp/gord-temp --profile=6061 --rpcmaxwebsockets=$NUM_CLIENTS &
GORD_PID=$!
GORD_KILLED=0
function killGordIfNotKilled() {
  if [ $KGORD_KILLED -eq 0 ]; then
    kill $GORD_PID
  fi
}
trap "killGordIfNotKilled" EXIT

sleep 1

rpc-idle-clients --devnet --profile=7000 -n=$NUM_CLIENTS
TEST_EXIT_CODE=$?

kill $GORD_PID

wait $GORD_PID
GORD_EXIT_CODE=$?
GORD_KILLED=1

echo "Exit code: $TEST_EXIT_CODE"
echo "Gord exit code: $GORD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $GORD_EXIT_CODE -eq 0 ]; then
  echo "rpc-idle-clients test: PASSED"
  exit 0
fi
echo "rpc-idle-clients test: FAILED"
exit 1
