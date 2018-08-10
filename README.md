# hellofabric

基于 [hyperledger/sample](https://github.com/hyperledger/fabric-samples) basic-network 的
[hyperledger/fabric-java-sdk](https://github.com/hyperledger/fabric-sdk-java) 示例。
参考了[这篇博文](https://medium.com/@lkolisko/hyperledger-fabric-sdk-java-basics-tutorial-a67b2b898410)。

可能不适用于生产环境。

运行环境：

| **fabric** | **v1.1.0** |
| --- | --- |
| **fabric-java-sdk** | **v1.1.0** |

## 安装

* 服务端: 请参考 [hyperledger/fabric prerequisites](http://hyperledger-fabric.readthedocs.io/en/release-1.1/prereqs.html)
* 客户端: 请参考 [hyperledger/fabric-java-sdk](https://github.com/hyperledger/fabric-sdk-java)

## 运行

* 服务端

```bash
# 启动网路
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

* 服务端 cli 方式测试链码

```bash
# 调用链码，读写
# set `a' to `123'
./invoke simple set '"a","123"'
# Expected result:
#   ...
#   [chaincodeCmd] chaincodeInvokeOrQuery -> INFO 063 Chaincode invoke successful. result: status:200
#   ...

# 调用链码，只读
# query value of `a'
./query simple get '"a"'
# Expected result:
#   ...
#   Query Result: 123
```

* Java 客户端

在 IDE 中运行 `src/main/java/Main.java`.
详见 [fabric-java-sdk-demofabric-java-sdk-demo 下的 README.md](fabric-java-sdk-demo/README.md).

## 链码

* 链码 `simple`:
    - `set(2)`: 将指定值写入对应的 key
    - `get(1)`: 查询指定 key 对应的值

* 链码调试以及网络起停：详见 [basic-network 下的 README.md](basic-network/README.md).
* 链码和通道配置在[这里](basic-network/fabric.conf), 更新配置后需要执行 `./generate.sh` 并重启网络后生效。

## Copyleft

* [Jack](https://github.com/QwertyJack)

HAVE FUN !!!
