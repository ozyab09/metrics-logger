#include "gauge.h"

namespace MetricsManager {

    Gauge::Gauge(std::string label, const double value) : BaseMetric(std::move(label)), Value{value} {}

    void Gauge::Increment() {
        Increment(1.0);
    }

    void Gauge::Increment(const double value) {
        Change(value);
    }

    void Gauge::Decrement() {
        Decrement(1.0);
    }

    void Gauge::Decrement(const double value) {
        Change(-1.0 * value);
    }

    void Gauge::Set(const double value) {
        Value.store(value);
    }

    void Gauge::Change(const double value) {
        auto current = Value.load();
        while (!Value.compare_exchange_weak(current, current + value)) {
        }
    }

    double Gauge::GetValue() const {
        return Value.load();
    }

    std::string Gauge::SerializeInPrometheus() {
        return Label + " " + std::to_string(Value.load()) + "\n";
    }

    web::json::value Gauge::SerializeInJson() {
        web::json::value result;
        result[Label] = Value.load();
        return result;
    }

}  // namespace prometheus