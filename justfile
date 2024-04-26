default:
  @just --list

set dotenv-load
set fallback

# build a local geth binary
build:
    make geth

# initialize a local geth node
init:
    ./build/bin/geth --db.engine pebble \
    --state.scheme=path \
    init geth-genesis-local.json

# start a local geth node
run: init
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

[macos]
clean:
    rm -rf ~/Library/Ethereum/

[linux]
clean:
    rm -rf ~/.ethereum/

clean-restart: clean run
