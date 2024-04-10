#pragma once

#include "base_metric.h"
#include <atomic>
#include <utility>

namespace MetricsManager {
    class Gauge : BaseMetric {
    public:
        Gauge(std::string label) : BaseMetric(std::move(label)) {};

        explicit Gauge(std::string label, double value);

        void Increment();

        void Increment(double);

        void Decrement();

        void Decrement(double);

        void Set(double);

        double GetValue() const;

        std::string SerializeInPrometheus() override;

        web::json::value SerializeInJson() override;

    private:
        void Change(double);
    private:
        std::atomic<double> Value{0.0};
    };
}