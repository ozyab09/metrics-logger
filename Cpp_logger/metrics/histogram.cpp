#include "histogram.h"

#include <utility>

namespace MetricsManager {
    Histogram::Histogram(const std::string& label, const std::vector<double>& buckets): BaseMetric(std::move(label)), Buckets(buckets) {
        std::sort(Buckets.begin(), Buckets.end());
        for (double bucket : Buckets) {
            BucketCounts[bucket] = 0;
        }
        BucketCounts[std::numeric_limits<double>::infinity()] = 0;
    }


    void Histogram::Observe(double value) {
        std::unique_lock<std::mutex> lock(Mutex);
        for (auto &bucket: BucketCounts) {
            if (value <= bucket.first) {
                bucket.second++;
            }
        }
        Sum += value;
        Count++;
    }

    double Histogram::GetSum() const {
        std::unique_lock<std::mutex> lock(Mutex);
        return Sum;
    }

    size_t Histogram::GetCount() const {
        std::unique_lock<std::mutex> lock(Mutex);
        return Count;
    }

    Histogram::DeserializedResult Histogram::GetValue() const {
        std::unique_lock<std::mutex> lock(Mutex);
        DeserializedResult result;
        result.BucketCounts = BucketCounts;
        result.Sum = Sum;
        result.Count = Count;
        return result;
    }

    std::string Histogram::SerializeInPrometheus() {
        DeserializedResult value = GetValue();

        std::ostringstream output;

        for (const auto &bucket: value.BucketCounts) {
            output << Label << "_bucket{le=\"" << bucket.first << "\"} " << bucket.second << "\n";
        }
        output << Label << "_sum " << std::fixed << std::setprecision(6) << value.Sum << "\n";
        output << Label << "_count " << value.Count << "\n";

        return output.str();
    }

    web::json::value Histogram::SerializeInJson() {
        web::json::value result;
        DeserializedResult value = GetValue();
        for (const auto &bucket: BucketCounts) {
            result[Label + "_bucket{le=\"" + std::to_string(bucket.first) + "\"}"] = bucket.second;
        }
        result[Label + "_sum"] = value.Sum;
        result[Label + "_count"] = value.Count;
        return result;
    }
}