#!/bin/bash
#
# Copyright IBM Corp All Rights Reserved
#
# SPDX-License-Identifier: Apache-2.0
#
# Exit on first error
set -e
. ./fabric.conf

# launch network; create channel and join peer to channel

# Now launch the CLI container in order to install, instantiate chaincode
# and prime the ledger with our 10 cars

for i in ${!CHANNEL[@]}
do
    cc=$i
    ch=${CHANNEL[$i]}

    for org in ${ORGS[@]}
    do
        for peer in ${PEERS[@]}
        do
            install_cc $cc $ch $org $peer
        done
    done

    instantiate_cc $cc $ch ${ORGS[0]} ${PEERS[0]}
done

