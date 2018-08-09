import org.apache.log4j.Logger;
import org.hyperledger.fabric.sdk.BlockEvent.TransactionEvent;
import org.hyperledger.fabric.sdk.*;
import org.hyperledger.fabric.sdk.exception.InvalidArgumentException;
import org.hyperledger.fabric.sdk.exception.ProposalException;
import org.hyperledger.fabric.sdk.exception.TransactionException;
import org.hyperledger.fabric.sdk.security.CryptoSuite;

import java.io.InputStream;
import java.util.Collection;
import java.util.HashMap;
import java.util.Map;
import java.util.Properties;
import java.util.concurrent.CompletableFuture;

/**
 * Basic functions
 *
 * @author jack
 *
 */
public class Util {
    public static final Logger log = Logger.getLogger(Util.class);
    public static final Properties properties;
    public static HFClient client;
    public static Map<String, Channel> channel = new HashMap<String, Channel>();

    private Util() {
    }

    static {
        properties = new Properties();
        try {

            // init config
            InputStream in = Util.class.getClassLoader().getResourceAsStream("demo.properties");
            properties.load(in);

            // init HF client
            CryptoSuite cryptoSuite = CryptoSuite.Factory.getCryptoSuite();
            client = HFClient.createNewInstance();
            client.setCryptoSuite(cryptoSuite);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    /**
     * Get or initialize the channel
     * <br>
     * the client must be bind with a user context first via `client.setUserContext(User)`.
     *
     * @param chan Channel
     * @return Initialized channel
     */
    public static Channel getChannel(String chan) throws InvalidArgumentException, TransactionException {
        if (!channel.containsKey(chan))
            return initChannel(chan);
        else
            return channel.get(chan);
    }

    /**
     * Force refresh the channel
     * <br>
     * the client must be bind with a user context first via `client.setUserContext(User)`.
     *
     * @param chan Channel
     * @return Initialized channel
     */
    public static Channel initChannel(String chan) throws InvalidArgumentException, TransactionException
    {
        Peer peer = client.newPeer(
                Util.properties.getProperty("peer"),
                Util.properties.getProperty("peerEndpoint"));
        EventHub eventHub = client.newEventHub("eventHub@" + chan,
                Util.properties.getProperty("eventHubEndpoint"));
        Orderer orderer = client.newOrderer(
                Util.properties.getProperty("orderer"),
                Util.properties.getProperty("ordererEndpoint"));
        Channel ch = client.newChannel(chan);

        ch.addPeer(peer);
        ch.addEventHub(eventHub);
        ch.addOrderer(orderer);
        ch.initialize();
        channel.put(chan, ch);
        return ch;
    }

    /**
     * Query chaincode, do not write
     *
     * @param ch Channel
     * @param cc ChainCode
     * @param fn Function
     * @param args Args
     * @return Response
     */
    public static String query(String ch, String cc, String fn, String[] args) throws InvalidArgumentException, ProposalException, TransactionException {
        QueryByChaincodeRequest qpr = client.newQueryProposalRequest();
        qpr.setChaincodeID(ChaincodeID.newBuilder().setName(cc).build());
        qpr.setFcn(fn);
        qpr.setArgs(args);
        Collection<ProposalResponse> res = getChannel(ch).queryByChaincode(qpr);
        return new String(res.iterator().next().getChaincodeActionResponsePayload());
    }

    /**
     * Invoke chaincode synchronously
     *
     * @param ch Channel
     * @param cc ChainCode
     * @param fn Function
     * @param args Args
     * @param txid Return the transaction ID
     * @return successful or not
     */
    public static boolean invoke(String ch, String cc, String fn, String[] args, StringBuilder txid) throws ProposalException, InvalidArgumentException, TransactionException {
        TransactionProposalRequest tpr = client.newTransactionProposalRequest();
        tpr.setChaincodeID(ChaincodeID.newBuilder().setName(cc).build());
        tpr.setFcn(fn);
        tpr.setArgs(args);
        Collection<ProposalResponse> responses = getChannel(ch).sendTransactionProposal(tpr);
        CompletableFuture<TransactionEvent> future = getChannel(ch).sendTransaction(responses);
        TransactionEvent event = future.join();
        txid.setLength(0);
        txid.append(event.getTransactionID());
        return event.isValid();
    }

    /**
     * Invoke chaincode asynchronously
     * <br>
     * use `.get` to get the result
     *
     * @param ch Channel
     * @param cc ChainCode
     * @param fn Function
     * @param args Args
     * @return The future of transaction event
     */
    public CompletableFuture<TransactionEvent> invoke(String ch, String cc, String fn, String[] args) throws ProposalException, InvalidArgumentException, TransactionException {
        TransactionProposalRequest tpr = client.newTransactionProposalRequest();
        tpr.setChaincodeID(ChaincodeID.newBuilder().setName(cc).build());
        tpr.setFcn(fn);
        tpr.setArgs(args);
        Collection<ProposalResponse> responses = getChannel(ch).sendTransactionProposal(tpr);
        return getChannel(ch).sendTransaction(responses);
    }
}
