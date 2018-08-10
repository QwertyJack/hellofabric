import org.hyperledger.fabric.sdk.exception.InvalidArgumentException;

public class Main {
    public static void main(String[] args) throws InvalidArgumentException {
        StringBuilder txid;
        String res;

        // set user context
        Util.client.setUserContext(UserManagement.getOrCreateUser("User1"));

        try {
            Util.log.info("Set value");
            txid = new StringBuilder();
            if (Util.invoke("mychannel", "simplecc", "set",
                    new String[]{"a", "123"}, txid))
                Util.log.info(txid);

            Util.log.info("Get value");
            res = Util.query("mychannel", "simplecc", "get", new String[]{"a"});
            Util.log.info(res);
        } catch (Exception e) {
            Util.log.error(e.getMessage());
        }
    }
}
