#!/bin/bash
ADDRESS=$1

# Start
/app/bin/monacoind -conf=/root/.monacoin/monacoin.conf # -server -rpcbind=0.0.0.0 -rpcallowip=0.0.0.0/0 -rpcuser=user -rpcpassword=password
sleep 10

# Print setup
echo "MONACOIN_ADDRESS=$ADDRESS"

# Import the address
/app/bin/monacoin-cli importaddress $ADDRESS

# Generate enough block to pass the maturation time
/app/bin/monacoin-cli generatetoaddress 101 $ADDRESS

# Simulate mining
while :
do
    /app/bin/monacoin-cli generatetoaddress 1 $ADDRESS
    sleep 10
done
