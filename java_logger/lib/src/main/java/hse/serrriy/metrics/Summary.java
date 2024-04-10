package hse.serrriy.metrics;

import java.util.ArrayList;

public class Summary {
    RingDoubleBuffer buffer;
    ArrayList<Double> percentileValues;

    Summary(int capacity, ArrayList<Double> percentiles) {
        buffer = new RingDoubleBuffer(capacity);
        this.percentileValues = percentiles;
    }

    private int getSize() {
        return buffer.size();
    }

    private double getSum() {
        return buffer.sum();
    }

    public void add(double value) {
        buffer.add(value);
    }

    synchronized SummaryPercentileResult getPercentiles() {
        ArrayList<Double> results = new ArrayList<>();
        for (double percentile : percentileValues) {
            results.add(buffer.getPercentile(percentile));
        }
        return new SummaryPercentileResult(
            percentileValues,
            results,
            buffer.sum(),
            buffer.size()
        );
    }
}

class SummaryPercentileResult {
    ArrayList<Double> percentileValues;
    ArrayList<Double> result;
    double sum;
    int count;

    SummaryPercentileResult(ArrayList<Double> percentileValues, ArrayList<Double> result, double sum, int count) {
        this.percentileValues = percentileValues;
        this.result = result;
        this.sum = sum;
        this.count = count;
    }

    String toString(String metricName) {
        StringBuilder stringBuilder = new StringBuilder();
        for (int i = 0; i < percentileValues.size(); i++) {
            String str = "%s{quantile=%s} %s\n".formatted(metricName, percentileValues.get(i), result.get(i));
            stringBuilder.append(str);
        }
        return stringBuilder.toString();
    }
}
