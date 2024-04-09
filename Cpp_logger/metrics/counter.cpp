#include "counter.h"

namespace MetricsManager {
    void Counter::Increment() {
        Increment(1.0);
    }

    void Counter::Increment(const double value) {
        Value.Increment(value);
    }

    void Counter::Reset() {
        Value.Set(0.0);
    }

    double Counter::GetValue() const {
        return Value.GetValue();
    }

    std::string Counter::SerializeInPrometheus() {
        return Label + " " + std::to_string(Value.GetValue()) + "\n";
    }

    web::json::value Counter::SerializeInJson() {
        web::json::value result;
        result[Label] = Value.GetValue();
        return result;
    }

}  // namespace prometheus