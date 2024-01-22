import com.google.gson.Gson;

public class Logger {
    private static final Gson gson = new Gson();

    public static void d(Object obj) {
        report(Level.DEBUG, obj);
    }

    public static void v(Object obj) {
        report(Level.VERBOSE, obj);
    }

    public static void i(Object obj) {
        report(Level.INFO, obj);
    }

    public static void w(Object obj) {
        report(Level.WARNING, obj);
    }

    public static void e(Object obj) {
        report(Level.ERROR, obj);
    }

    private static void report(Level level, Object obj) {
        // TODO clear way to text
        System.out.println(level.name() + " " + gson.toJson(obj));
    }
}

