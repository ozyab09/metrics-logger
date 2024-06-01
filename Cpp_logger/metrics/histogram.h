#pragma once

#include <map>
#include <iomanip>
#include <atomic>
#include <mutex>
#include <algorithm>

#include "base_metric.h"

namespace MetricsManager {
    class Histogram : BaseMetric {
    public:
        Histogram(const std::string& label, const std::vector<double>& buckets);

        struct DeserializedResult {
            std::map<double, int> BucketCounts;
            double Sum;
            size_t Count;
        };

        double GetSum() const;

        size_t GetCount() const;

        DeserializedResult GetValue() const;

        void Observe(double value);

        std::string SerializeInPrometheus() override;

        web::json::value SerializeInJson() override;

    private:
        std::vector<double> Buckets;
        std::map<double, int> BucketCounts;
        double Sum = 0.0;
        int Count = 0;
        mutable std::mutex Mutex;
    };
}