# Fabric-java-sdk-demo

## IDE setup

* 确保 IDE 支持 Maven, 并且可以联网更新 Maven 依赖
* 先检出项目，也可以使用 IDE 自带的 VCS.
* Java demo 根目录是 `/fabric-java-sdk-demo`
* 主程序入口为 `/fabric-java-sdk/demo/src/main/java/Main.java`

以下测试通过:

* Eclipse: mars 及以上版本
    - "文件" -> "导入" -> "Existing Maven Projects" -> 选 Java demo 根目录
    - 运行主程序

* Idea: 2017 及以上版本
    - 开业界面 -> "打开" -> 选 Java demo 根目录
    - 打开后更新 Maven
    - 运行主程序

## 运行 demo

1. `/basic-network` 下执行 `./restart`
2. 根据实际情况修改资源目录下的配置文件，即 `src/main/resources/demo.properties`
   推荐使用 ssh 端口转发，将运行 basic-network 的服务器对应 endpoint 端口转发到本地
3. 运行项目目录下的`src/main/java/Main.java`

## 常见问题

* fabric 版本 1.1.0/1.0.0, fabric-java-sdk 1.1.0/1.0.0 均测试通过；1.2.0 会报错
* IDE 中项目根目录应选择`/fabric-java-sdk-demo`而不是 git 仓库根目录
* `fabric-java-sdk` 配置文件位于项目根目录，即`/fabric-java-sdk-demo`，并不是资源目录 `/fabric-java-sdk-demo/src/main/resources`
* 如果编译失败请检查 Maven 依赖是否正确安装
* 在 basic-network 加密套件保持不变的情况下，admin密钥文件保持不变，此密钥只允许注册一次，如果误删除请重启ca
* 测试时报用户密钥无效，请删除项目根目录下所有`.jso`文件；进一步，考虑重启网络

