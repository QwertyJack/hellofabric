# hellofabric

This is a toy demo of [hyperledger/fabric](https://github.com/hyperledger/fabric) based on
`basic-network` at [hyperledger/sample](https://github.com/hyperledger/fabric-samples), featured with
[hyperledger/fabric-java-sdk](https://github.com/hyperledger/fabric-sdk-java).

This demo highly refers to [this blog](https://medium.com/@lkolisko/hyperledger-fabric-sdk-java-basics-tutorial-a67b2b898410).

Just a demo, not for production env.

Tested under:

* fabric v1.1.0
* fabric-java-sdk v1.1.0

## Install Fabric

* server: see [hyperledger/fabric prerequisites](http://hyperledger-fabric.readthedocs.io/en/release-1.1/prereqs.html)
* client: see [hyperledger/fabric-java-sdk](https://github.com/hyperledger/fabric-sdk-java)

## IDE support

* IDEA
* Eclipse

## Run the demo

* server

```bash
# start the network
cd basic-network
./restart
# Expected result:
#   ...
#   Creating cli ... done
#   Install chaincode: simple, version: 1.0
#   ...
#   [chaincodeCmd] install -> DEBU 00e Installed remotely response:<status:200 payload:"OK" >
#   [main] main -> INFO 00f Exiting.....
#   Instantiate chaincode: simple, channel: mychannel
#   ...
#   [msp/identity] Sign -> DEBU 008 Sign: plaintext: ...
#   [msp/identity] Sign -> DEBU 009 Sign: digest: ...
#   [main] main -> INFO 00a Exiting.....
```

* cli client

```bash
# invoke chaincode
# set `a' to `123'
./invoke simple set '"a","123"'
# Expected result:
#   ...
#   [chaincodeCmd] chaincodeInvokeOrQuery -> INFO 063 Chaincode invoke successful. result: status:200
#   ...

# query chaincode
# query value of `a'
./query simple get '"a"'
# Expected result:
#   ...
#   Query Result: 123
```

* java client

Open project in your IDE and run `src/Main.java`.
See also [README.md under fabric-java-sdk-demofabric-java-sdk-demo](fabric-java-sdk-demo/README.md).

## Copyleft

* [Jack](https://github.com/QwertyJack)

HAVE FUN !!!
