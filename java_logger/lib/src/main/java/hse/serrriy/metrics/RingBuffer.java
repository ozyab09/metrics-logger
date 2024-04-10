package hse.serrriy.metrics;

import java.util.ArrayDeque;
import java.util.Deque;
import java.util.List;

class RingBuffer<T> {
    private final int DEFAULT_CAPACITY = 1028;

    protected final Deque<T> values;
    private final int capacity;

    RingBuffer(int capacity) {
        this.capacity = capacity;
        values = new ArrayDeque<>();
    }

    RingBuffer() {
        this.capacity = DEFAULT_CAPACITY;
        values = new ArrayDeque<>();
    }

    synchronized void add(T value) {
        values.addLast(value);
        if (values.size() > capacity) {
            values.removeFirst();
        }
    }

    synchronized T getPercentile(double percentile) {
        if (values.isEmpty()) {
            throw new IllegalArgumentException("The input dataset cannot be empty.");
        }
        if (percentile < 0 || percentile > 100) {
            throw new IllegalArgumentException("Percentile must be between 0 and 100 inclusive.");
        }
        List<T> sortedList = values.stream()
                .sorted()
                .toList();

        int rank = percentile == 0 ? 1 : (int) Math.ceil(percentile / 100.0 * values.size());
        return sortedList.get(rank - 1);
    }

    synchronized int size() {
        return values.size();
    }
}

class RingDoubleBuffer extends RingBuffer<Double> {

    RingDoubleBuffer(int capacity) {
        super(capacity);
    }

    synchronized double sum() {
        double result = 0.;
        for (double value : super.values) {
            result += value;
        }
        return result;
    }
}
