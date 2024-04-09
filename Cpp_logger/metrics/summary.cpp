#include "summary.h"

#include <utility>

namespace MetricsManager {
    Summary::Summary(std::string label,
                     const std::vector<double>& percentiles,
                     size_t chunkSize) : BaseMetric(std::move(label)),
                                         Percentiles(percentiles),
                                         Values(tag::tail<right>::cache_size = chunkSize) {}


    void Summary::Observe(double value) {
        std::unique_lock<std::mutex> lock(Mutex);
        Values(value);
    }

    double Summary::GetSum() const {
        std::unique_lock<std::mutex> lock(Mutex);
        return boost::accumulators::sum(Values);
    }

    size_t Summary::GetCount() const {
        std::unique_lock<std::mutex> lock(Mutex);
        return boost::accumulators::count(Values);
    }

    Summary::DeserializedResult Summary::GetValue() const {
        std::unique_lock<std::mutex> lock(Mutex);
        DeserializedResult result;
        result.Value.reserve(Percentiles.size());
        for (const auto& q : Percentiles) {
            result.Value.emplace_back(q, boost::accumulators::quantile(Values, quantile_probability = q));
        }
        result.Sum = boost::accumulators::sum(Values);
        result.Count = boost::accumulators::count(Values);
        return result;
    }

    std::pair<double, double> Summary::GetValueByPercentile(double percentile) const {
        return GetValueByPercentile(std::vector({percentile})).front();
    }

    std::vector<std::pair<double, double>> Summary::GetValueByPercentile(const std::vector<double>& percentiles) const {
        std::unique_lock<std::mutex> lock(Mutex);
        std::vector<std::pair<double, double>> result;
        result.reserve(percentiles.size());
        for (const auto& q : percentiles) {
            result.emplace_back(q, boost::accumulators::quantile(Values, quantile_probability = q));
        }
        return result;
    }

    std::string Summary::SerializeInPrometheus() {
        std::string result;
        DeserializedResult value = GetValue();
        for (const auto& it : value.Value) {
            result += Label + "{quantile=\"" + std::to_string(it.first) + "\"} " + std::to_string(it.second) + "\n";
        }
        result += Label + "_sum " + std::to_string(value.Sum) + "\n";
        result += Label + "_count " + std::to_string(value.Count) + "\n";
        return result;
    }

    web::json::value Summary::SerializeInJson() {
        web::json::value result;
        DeserializedResult value = GetValue();
        for (const auto& it : value.Value) {
            result[Label + "{quantile=\"" + std::to_string(it.first) + "\"}"] = it.second;
        }
        result[Label + "_sum"] = value.Sum;
        result[Label + "_count"] = value.Count;
        return result;
    }
}