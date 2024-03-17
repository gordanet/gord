#!/bin/bash

APPDIR=/tmp/gord-temp
GORD_RPC_PORT=29587

rm -rf "${APPDIR}"

gord --simnet --appdir="${APPDIR}" --rpclisten=0.0.0.0:"${GORD_RPC_PORT}" --profile=6061 &
GORD_PID=$!

sleep 1

RUN_STABILITY_TESTS=true go test ../ -v -timeout 86400s -- --rpc-address=127.0.0.1:"${GORD_RPC_PORT}" --profile=7000
TEST_EXIT_CODE=$?

kill $GORD_PID

wait $GORD_PID
GORD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Gordd exit code: $GORD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $GORD_EXIT_CODE -eq 0 ]; then
  echo "mempool-limits test: PASSED"
  exit 0
fi
echo "mempool-limits test: FAILED"
exit 1
