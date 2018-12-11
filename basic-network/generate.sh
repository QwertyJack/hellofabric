#!/bin/bash
#
# Copyright IBM Corp All Rights Reserved
#
# SPDX-License-Identifier: Apache-2.0
#
export PATH=$GOPATH/src/github.com/hyperledger/fabric/build/bin:~/bin:$PATH
export FABRIC_CFG_PATH=${PWD}
CHANNEL_NAME=mychannel

while getopts n opt
do
    case $opt in
        n)  RENEW=1;;
        ?)  ;;
    esac
done

if [[ x$RENEW = x1 ]]
then
    echo Re-generate crypto-marterial...

    # remember sk, by Jack
    sk_old=$(find crypto-config/peerOrganizations/org1.example.com/ca -name "*_sk" -exec basename {} \;)

    # remove previous crypto material and config transactions
    rm -fr config/*
    rm -fr crypto-config/*

    # generate crypto material
    cryptogen generate --config=./crypto-config.yaml
    if [ "$?" -ne 0 ]; then
        echo "Failed to generate crypto material..."
        exit 1
    fi

    # update sk, by Jack
    sk_new=$(find crypto-config/peerOrganizations/org1.example.com/ca -name "*_sk" -exec basename {} \;)
    sed -i "s/"$sk_old"/"$sk_new"/" ./docker-compose.yml
fi

echo Generate genesis block for orderer
configtxgen -profile OneOrgOrdererGenesis -outputBlock ./config/genesis.block
if [ "$?" -ne 0 ]; then
    echo "Failed to generate orderer genesis block..."
    exit 1
fi

. ./fabric.conf
for chan in `for i in ${CHANNEL[@]}; do echo $i; done | sort -u`
do
    gen_crypto_4_channel $chan
done
