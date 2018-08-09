#! /bin/sh
#
# func.sh
# Copyright (C) 2018 jack <jack@HP-WorkStation>
#
# Distributed under terms of the MIT license.
#

# adapt `docker.compse` for snap

docker-compose -v &>/dev/null || {
    shopt -s expand_aliases
    alias docker-compose=docker.compose
}

# util functions

gen_crypto_4_channel () {

    CHANNEL_NAME=$1
    echo Generate config for channel: $CHANNEL_NAME

    # generate channel configuration transaction
    configtxgen -profile OneOrgChannel -outputCreateChannelTx ./config/channel_$CHANNEL_NAME.tx -channelID $CHANNEL_NAME
    if [ "$?" -ne 0 ]; then
        echo "Failed to generate channel configuration transaction..."
        exit 1
    fi

    # generate anchor peer transaction
    configtxgen -profile OneOrgChannel -outputAnchorPeersUpdate ./config/Org1MSPanchors_$CHANNEL_NAME.tx -channelID $CHANNEL_NAME -asOrg Org1MSP
    if [ "$?" -ne 0 ]; then
        echo "Failed to generate anchor peer update for Org1MSP..."
        exit 1
    fi
}

create_and_join_channel () {
    CHANNEL_NAME=$1

    # Create the channel
    echo Create channel: $CHANNEL_NAME
    docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/msp/users/Admin@org1.example.com/msp" peer0.org1.example.com peer channel create -o orderer.example.com:7050 -c $CHANNEL_NAME -f /etc/hyperledger/configtx/channel_$CHANNEL_NAME.tx

    # Join peer0.org1.example.com to the channel.
    echo Join channel: $CHANNEL_NAME
    docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/msp/users/Admin@org1.example.com/msp" peer0.org1.example.com peer channel join -b $CHANNEL_NAME.block
}

install_and_instantiate() {
    echo Install chaincode: $1, version: $VERSION
    docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp" cli peer chaincode install -n "$1"cc -v "$VERSION" -p github.com/$1

    ARGS='{"Args":[""]}'
    echo Instantiate chaincode: $1, channel: $2
    docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp" cli peer chaincode instantiate -o orderer.example.com:7050 -C "$2" -n "$1"cc -v $VERSION -c "$ARGS" -P "$POLICY"
}

upgrade() {
    VERSION=$3

    docker cp ../chaincode/$1 cli:/opt/gopath/src/github.com

    docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp" cli peer chaincode install -n "$1"cc -v "$VERSION" -p github.com/$1

    ARGS='{"Args":[""]}'
    docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp" cli peer chaincode upgrade -o orderer.example.com:7050 -C "$2" -n "$1"cc -v $VERSION -c "$ARGS" -P "$POLICY"
}
