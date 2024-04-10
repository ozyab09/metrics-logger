package hse.serrriy;

import hse.serrriy.metrics.Metrics;

public class MetricsLoggerTest {
    public static void main(String[] args) {
        Config config = new Config("SOME_SERVICE", "logs", "metrics", 8080);

        Logger logger = new Logger();
        Metrics metrics = new Metrics();

        Reporter.start(config, logger, metrics);

        logger.d("log1");
        logger.d("log2");

        metrics.createCounter("counter_ex");
        metrics.getCounter("counter_ex").increment();

        metrics.createGauge("gauge_ex");
        metrics.getGauge("gauge_ex").set(45);

        metrics.createSummary("summary_ex", 10, 50);
        metrics.getSummary("summary_ex").add(1);
        metrics.getSummary("summary_ex").add(5);
        metrics.getSummary("summary_ex").add(-12);
        metrics.getSummary("summary_ex").add(100);
    }
}
