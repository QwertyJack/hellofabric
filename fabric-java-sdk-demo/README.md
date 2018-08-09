# Fabric-java-sdk-demo

## IDE setup

### Eclipse
1. Git 视图下克隆项目并检出
2. 打开终端切换到 `/fabric-java-sdk-demo`，执行 `mvn eclipse:eclipse`安装 `fabric-java-sdk`以及依赖
3. 回到 Eclipse 的 Git 视图，选择 `/fabric-java-sdk-demo`，右键导入项目
4. 导入后自动编译

### Idea
1. 用 Git 克隆项目并检出
2. 打开，目录选择`/fabric-java-sdk-demo`
3. 导入后第一次编译会自动安装 Maven 依赖

## Run test
1. 启动basic-network
2. 编译项目
3. fabric network 的 endpoint 定义在资源目录下的`src/demo.properties`，确保本地可以访问；推荐使用 ssh 本地端口转发将运行basic-network的服务器对应endpoint端口转发到本地
4. 示例文件位于项目根目录下的`src/Main.java`，运行示例请运行文件对应的`main`方法

## Troubleshooting
* 注意 IDE 中项目根目录应选择`/fabric-java-sdk-demo`而不是 Git 仓库根目录
* Eclipse 需要手动安装 Maven 依赖；当然也可以当作 Maven 项目导入并执行 Maven Install
* `fabric-java-sdk`配置文件位于项目根目录，即`/fabric-java-sdk-demo`，并不是资源目录
* 如果编译失败请检查 Maven 依赖是否正确安装
* 在basic-network加密套件保持不变的情况下，admin密钥文件保持不变，此密钥只允许注册一次，如果误删除请重启ca
* 测试时报用户密钥无效请删除项目根目录下所有`.jso`文件；进一步，考虑重启ca
* 账本为空时所有测试通过，如果账本不为空测试可能失败；清空账本最简单的方法是重启basic-network

