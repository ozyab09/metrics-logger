#pragma once

#include <string>
#include <utility>
#include <cpprest/json.h>

namespace MetricsManager {
    class BaseMetric {
    public:
        virtual std::string SerializeInPrometheus() = 0;
        virtual web::json::value SerializeInJson() = 0;

        BaseMetric(std::string label) : Label(std::move(label)) {};
    public:
        std::string Label;
    };
}