package hse.serrriy.metrics;

public class Counter {
    private double value = 0;

    synchronized double getValue() {
        return value;
    }

    public synchronized void increment() {
        increment(1);
    }

    public synchronized void increment(int value) {
        this.value += value;
    }

    String toString(String metricName) {
        return "%s %s".formatted(metricName, getValue());
    }
}
