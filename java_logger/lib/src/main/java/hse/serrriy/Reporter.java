package hse.serrriy;

import hse.serrriy.metrics.Metrics;
import io.javalin.Javalin;

public class Reporter {
    public static void start(Config config, Logger logger, Metrics metrics) {
        Javalin app = Javalin.create()
                .start(config.port);

        app.get("/%s".formatted(config.logsPath), ctx -> {
            ctx.result(logger.logs.toString());
        });
        app.get("/%s".formatted(config.metricsPath), ctx -> {
            ctx.result(metrics.getAll());
        });
    }
}
