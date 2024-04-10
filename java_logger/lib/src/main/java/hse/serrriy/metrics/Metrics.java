package hse.serrriy.metrics;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.function.BiConsumer;

public class Metrics {
    private final HashMap<String, Counter> counters = new HashMap<>();
    private final HashMap<String, Gauge> gauges = new HashMap<>();
    private final HashMap<String, Summary> summaries = new HashMap<>();

    public void createCounter(String name) {
        counters.putIfAbsent(name, new Counter());
    }

    public Counter getCounter(String name) {
        Counter counter = counters.get(name);
        if (counter == null) {
            throw new IllegalArgumentException("Create Counter metric first");
        }
        return counter;
    }

    public void createGauge(String name) {
        gauges.putIfAbsent(name, new Gauge());
    }

    public Gauge getGauge(String name) {
        Gauge gauge = gauges.get(name);
        if (gauge == null) {
            throw new IllegalArgumentException("Create Gauge metric first");
        }
        return gauge;
    }

    public void createSummary(String name, int capacity, double... percentiles) {
        ArrayList<Double> percentilesArray = new ArrayList<>();
        for (double percentile : percentiles) {
            percentilesArray.add(percentile);
        }
        summaries.putIfAbsent(name, new Summary(capacity, percentilesArray));
    }

    public Summary getSummary(String name) throws IllegalArgumentException {
        Summary summary = summaries.get(name);
        if (summary == null) {
            throw new IllegalArgumentException("Create Summary metric first");
        }
        return summary;
    }

    public String getAll() {
        StringBuilder result = new StringBuilder();
        counters.forEach((name, counter) -> {
            result.append(counter.toString(name));
        });
        result.append("\n");
        gauges.forEach((name, gauge) -> {
            result.append(gauge.toString(name));
        });
        result.append("\n");
        summaries.forEach((name, summary) -> {
            result.append(summary.getPercentiles().toString(name));
        });
        return result.toString();
    }
}
