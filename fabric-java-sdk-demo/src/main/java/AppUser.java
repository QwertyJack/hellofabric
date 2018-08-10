import org.hyperledger.fabric.sdk.Enrollment;
import org.hyperledger.fabric.sdk.User;

import java.io.IOException;
import java.io.ObjectInputStream;
import java.io.ObjectOutputStream;
import java.io.Serializable;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.Set;

/**
 * Basic implementation of the {@link User} interface.
 *
 * @author jack
 *
 */
public class AppUser implements User, Serializable {

    private static final long serialVersionUID = 1L;
    private String name;
    private Set<String> roles;
    private String account;
    private String affiliation;
    private Enrollment enrollment;
    private String mspId;

    public AppUser(String name, String affiliation, String mspId, Enrollment enrollment) {
        this.name = name;
        this.affiliation = affiliation;
        this.enrollment = enrollment;
        this.mspId = mspId;
    }

    public String getName() {
        return name;
    }

    public Set<String> getRoles() {
        return roles;
    }

    public String getAccount() {
        return account;
    }

    public String getAffiliation() {
        return affiliation;
    }

    public Enrollment getEnrollment() {
        return enrollment;
    }

    public String getMspId() {
        return mspId;
    }

    @Override
    public String toString() {
        return "AppUser: " + name + "\n" + enrollment.getCert();
    }

    /**
     * Load user from local binary; if failed, then create
     *
     * @param name User name
     * @return The user
     */
    public static AppUser load(String name) throws IOException, ClassNotFoundException {
        if (Files.exists(Paths.get(name + ".jso"))) {
             ObjectInputStream decoder = new ObjectInputStream(
                     Files.newInputStream(Paths.get(name + ".jso")));
             return (AppUser)decoder.readObject();
        } else {
            return null;
        }
    }

    /**
     * Save user obj to local
     */
    public void save() throws IOException {
        ObjectOutputStream oos = new ObjectOutputStream(
                Files.newOutputStream(Paths.get(this.name + ".jso")));
        oos.writeObject(this);
    }
}