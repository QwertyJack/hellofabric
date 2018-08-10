import java.lang.reflect.InvocationTargetException;
import java.net.MalformedURLException;

import org.hyperledger.fabric.sdk.Enrollment;
import org.hyperledger.fabric.sdk.exception.CryptoException;
import org.hyperledger.fabric.sdk.security.CryptoSuite;
import org.hyperledger.fabric_ca.sdk.HFCAClient;
import org.hyperledger.fabric_ca.sdk.RegistrationRequest;

/**
 * User load/enroll
 *
 * @author jack
 *
 */
public class UserManagement {

    private static HFCAClient caClient = null;

    private static void initCAClient() throws MalformedURLException, IllegalAccessException, InstantiationException, ClassNotFoundException, CryptoException, org.hyperledger.fabric.sdk.exception.InvalidArgumentException, NoSuchMethodException, InvocationTargetException {
        // build CA client
        CryptoSuite cryptoSuite = CryptoSuite.Factory.getCryptoSuite();
        caClient = HFCAClient.createNewInstance(
                Util.properties.getProperty("caEndpoint"), null);
        caClient.setCryptoSuite(cryptoSuite);
    }

    /**
     * Load or create admin
     *
     * @return An admin
     */
    public static AppUser getOrCreateAdmin() {
        AppUser admin = null;
        try {
            admin = AppUser.load(
                    Util.properties.getProperty("admin"));
            if (admin == null) {
                if (null == caClient)
                    initCAClient();
                Util.log.debug("Enrolling admin ...");
                Enrollment adminEnrollment = caClient.enroll(
                        Util.properties.getProperty("admin"),
                        Util.properties.getProperty("admin_password"));
                admin = new AppUser(
                        Util.properties.getProperty("admin"),
                        Util.properties.getProperty("affiliation"),
                        Util.properties.getProperty("mspId"), adminEnrollment);
                admin.save();
            }
        } catch (Exception e) {
            // TODO Auto-generated catch block
            e.printStackTrace();
        }

        return admin;
    }

    /**
     * Get or create an user by name
     *
     * @param userId User name
     * @return An user
     */
    public static AppUser getOrCreateUser(String userId) {
        AppUser appUser = null;
        try {
            appUser = AppUser.load(userId);
            if (appUser == null) {
                Util.log.debug("Enrolling user: " + userId);
                if (Util.properties.getProperty("admin").equals(userId))
                    return getOrCreateAdmin();
                if (null == caClient)
                    initCAClient();
                RegistrationRequest rr = new RegistrationRequest(userId,
                        Util.properties.getProperty("affiliation"));
                String enrollmentSecret = caClient.register(rr, getOrCreateAdmin());
                Enrollment enrollment = caClient.enroll(userId, enrollmentSecret);
                appUser = new AppUser(userId,
                        Util.properties.getProperty("affiliation"),
                        Util.properties.getProperty("mspId"), enrollment);
                appUser.save();
            }
        } catch (Exception e) {
            // TODO Auto-generated catch block
            e.printStackTrace();
        }

        return appUser;
    }
}
