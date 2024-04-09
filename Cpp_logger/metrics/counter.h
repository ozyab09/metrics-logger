#pragma once

#include <utility>

#include "gauge.h"

namespace MetricsManager {
    class Counter : BaseMetric {
    public:
        Counter(std::string label) : BaseMetric(std::move(label)) {};

        void Increment();

        void Increment(double);

        void Reset();

        double GetValue() const;

        std::string SerializeInPrometheus() override;

        web::json::value SerializeInJson() override;

    private:
        Gauge Value{"", 0.0};
    };
}