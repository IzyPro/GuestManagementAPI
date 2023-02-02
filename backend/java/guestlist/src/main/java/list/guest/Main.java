package list.guest;

import java.sql.Connection;
import java.sql.DriverManager;
import java.util.logging.Logger;

public class Main {
    public static void main(String[] args) {
        Connection conn = null;
        try {
            conn = DriverManager.getConnection(
                    "jdbc:mysql://mysql:3306/database?useSSL=false",
                    "user",
                    "password");
	    conn.close();
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}

