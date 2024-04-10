package hse.serrriy;

import com.google.gson.Gson;

import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.io.OutputStream;
import java.nio.charset.StandardCharsets;
import java.time.LocalDateTime;

public class Logger {
    Logger() {}

    OutputStream logs = new ByteArrayOutputStream();

    private final Gson gson = new Gson();

    public void d(Object obj) {
        report(Level.DEBUG, obj);
    }

    public void v(Object obj) {
        report(Level.VERBOSE, obj);
    }

    public void i(Object obj) {
        report(Level.INFO, obj);
    }

    public void w(Object obj) {
        report(Level.WARNING, obj);
    }

    public void e(Object obj) {
        report(Level.ERROR, obj);
    }

    private synchronized void report(Level level, Object obj) {
        String log = LocalDateTime.now() + " " + level.name() + " " + gson.toJson(obj) + " " + Thread.currentThread().getName() + System.lineSeparator();
        try {
            logs.write(log.getBytes(StandardCharsets.UTF_8));
        } catch (IOException e) {}
    }
}

