#!/bin/bash
#
# Copyright IBM Corp All Rights Reserved
#
# SPDX-License-Identifier: Apache-2.0
#
# Exit on first error, print all commands.
set -e
. ./fabric.conf

#docker-compose -f docker-compose.yml up -d ca.example.com orderer.example.com peer0.org1.example.com couchdb
docker-compose -f docker-compose.yml up -d

# wait for Hyperledger Fabric to start
# incase of errors when running later commands, issue export FABRIC_START_TIMEOUT=<larger number>
export FABRIC_START_TIMEOUT=10
#echo ${FABRIC_START_TIMEOUT}
#sleep ${FABRIC_START_TIMEOUT}

for chan in `for i in ${CHANNEL[@]}; do echo $i; done | sort -u`
do
    create_channel $chan
    for org in ${ORGS[@]}
    do
        for peer in ${PEERS[@]}
        do
            join_channel $chan $org $peer
        done
    done
done
