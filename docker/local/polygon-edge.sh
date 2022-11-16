#!/bin/sh

set -e

WOOP_CHAIN_BIN=./woopchain

case "$1" in

   "init")
      echo "Generating secrets..."
      secrets=$("$WOOP_CHAIN_BIN" secrets init --num 4 --data-dir data- --json)
      echo "Secrets have been successfully generated"

      echo "Generating genesis file..."
      "$WOOP_CHAIN_BIN" genesis \
        --dir /genesis/genesis.json \
        --consensus ibft \
        --ibft-validators-prefix-path data- \
        --bootnode /dns4/node-1/tcp/1478/p2p/$(echo $secrets | jq -r '.[0] | .node_id') \
        --bootnode /dns4/node-2/tcp/1478/p2p/$(echo $secrets | jq -r '.[1] | .node_id')
      echo "Genesis file has been successfully generated"
      ;;

   *)
      echo "Executing woopchain..."
      exec "$WOOP_CHAIN_BIN" "$@"
      ;;

esac
