### SOLO-1peer-testing Network Usage

* Edit `./fabric.conf` at your need: add chaincode and its channel under var `CHANNEL`.
* Run `./generate.sh` if new channel is added.
* To (re-)start the network, run `./restart.sh`;
  you may restart it every time when adding new chaincode or channel is reconfigured.
* Once chaincode is **installed successully**, you may want to upgrade the chaincode without restarting the network:

```bash
./upgrade.sh <chaincode_name> <version>
```

NOTE:
  * the chaincode is at `/chaincode/<chaincode_name>`
  * the `version` **must be different**, even if build failed; the init version is `1.0`

* Call chaincode via `./query` or `./invoke`

```bash
./(query | invoke) <chaincode_name> <function> <parameter>
```

The `parameter` must in the form of `'"par1", "par2", ...'`, i.e.

```bash
# call function `myfunction' of chaincode `test' with parameter `par1', ...
./query test myfunction '"par1", "par2", "par3"'
./invoke test myfunction '"par1"'
```

* `./teardown.sh` to shutdown.

### Chaincode Naming Spec

For example, chaincode `simple`:

* src: `/chaincode/simple`
* chaincode name: `simplecc`
