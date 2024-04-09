#pragma once

#include <atomic>
#include <mutex>

#include <boost/accumulators/accumulators.hpp>
#include <boost/accumulators/statistics.hpp>
#include "base_metric.h"

namespace MetricsManager {
    using namespace boost::accumulators;
    class Summary : BaseMetric {
    public:
        struct DeserializedResult {
            std::vector<std::pair<double, double>> Value;
            double Sum;
            size_t Count;
        };

        Summary(std::string label, const std::vector<double>& percentiles, size_t chunkSize = 1024);

        void Observe(double value);

        double GetSum() const;

        size_t GetCount() const;

        DeserializedResult GetValue() const;

        std::pair<double, double> GetValueByPercentile(double percentile) const;

        std::vector<std::pair<double, double>> GetValueByPercentile(const std::vector<double>& percentiles) const;

        std::string SerializeInPrometheus() override;

        web::json::value SerializeInJson() override;
    private:
        accumulator_set<double, stats<tag::sum, tag::count, tag::tail_quantile<right>>> Values;
        std::vector<double> Percentiles;
        mutable std::mutex Mutex;
    };
}