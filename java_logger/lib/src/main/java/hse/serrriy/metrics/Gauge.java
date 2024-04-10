package hse.serrriy.metrics;

import com.google.common.util.concurrent.AtomicDouble;

public class Gauge {
    private final AtomicDouble value = new AtomicDouble(0);

    synchronized double getValue() {
        return value.doubleValue();
    }

    public synchronized void increment() {
        add(1);
    }

    public synchronized void decrement() {
        decrement(1);
    }

    public synchronized void decrement(int value) {
        add(-1 * value);
    }

    public synchronized void add(int value) {
        this.value.addAndGet(value);
    }

    public synchronized void set(int value) {
        this.value.set(value);
    }

    String toString(String metricName) {
        return "%s %s".formatted(metricName, getValue());
    }
}
