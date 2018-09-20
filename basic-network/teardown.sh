#!/bin/bash
#
# Copyright IBM Corp All Rights Reserved
#
# SPDX-License-Identifier: Apache-2.0
#
# Exit on first error, print all commands.
set -e
. ./fabric.conf

# Shut down the Docker containers for the system tests.
docker-compose -f docker-compose.yml kill && docker-compose -f docker-compose.yml down

# remove the local state
rm -f ~/.hfc-key-store/*

# remove images, by Jack
docker rm $(docker ps -qa --filter name=dev-*)

# remove chaincode docker images
docker rmi $(docker images dev-* -q)

# Your system is now clean
