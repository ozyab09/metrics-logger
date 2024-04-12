#include "json.hpp"
#include <any>
#include <string>

#include "logger/logger_base.h"
#include "metrics/metrics.h"


int main() {
    MetricsLogger::BaseLogger logger;

    MetricsManager::MetricsManager metricsManager("config.json");
    metricsManager.RegisterCounter("CounterExample");
    metricsManager.RegisterGauge("GaugeExample");
    metricsManager.RegisterSummary("SummaryExample", {0.3, 0.5, 0.95, 0.99});
    for (int i = 0; i < 10000; ++i) {
        metricsManager.GetCounter("CounterExample")->Increment();
        metricsManager.GetGauge("GaugeExample")->Set(rand() % 100);
        metricsManager.GetSummary("SummaryExample")->Observe(i);
        logger.Log(MetricsLogger::LogLevel::WARN, "Metrics has been updated");
        sleep(1);
    }
    return 0;
}
