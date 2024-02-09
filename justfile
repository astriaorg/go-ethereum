default:
  @just --list

set dotenv-load
set fallback

build:
    make geth

init:
    ./build/bin/geth --db.engine pebble \
    --state.scheme=path \
    init geth-genesis-local.json

run:
    ./build/bin/geth \
    --networkid=1337 \
    --http \
    --http.addr=0.0.0.0 \
    --http.port=8545 \
    --http.corsdomain="*" \
    --http.vhosts="*" \
    --http.api=eth,net,web3,debug,txpool \
    --ws \
    --ws.addr=0.0.0.0 \
    --ws.port=8546 \
    --ws.origins="*" \
    --grpc \
    --grpc.addr=0.0.0.0 \
    --grpc.port=50051 \
    --db.engine=pebble \
    --state.scheme=path

# this is specifically for mac os
clean:
    rm -rf ~/Library/Ethereum/
