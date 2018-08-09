#!/bin/bash
#
# Copyright IBM Corp All Rights Reserved
#
# SPDX-License-Identifier: Apache-2.0
#
# Exit on first error
set -e
. ./fabric.conf

# don't rewrite paths for Windows Git Bash users
export MSYS_NO_PATHCONV=1

# launch network; create channel and join peer to channel

# Now launch the CLI container in order to install, instantiate chaincode
# and prime the ledger with our 10 cars

docker-compose -f ./docker-compose.yml up -d cli

for i in ${!CHANNEL[@]}
do
    install_and_instantiate $i ${CHANNEL[$i]}
done
