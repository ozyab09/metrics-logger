#include "json.hpp"
#include <any>
#include <string>

#include "logger/logger_base.h"
#include "metrics/metrics.h"


int main() {
    MetricsLogger::BaseLogger logger;
    std::any val = std::string("log value");
    logger.LogInFile(MetricsLogger::LogLevel::ERROR, "Log in console", "file.txt", {{"Optional log", val}});

    MetricsManager::MetricsManager metricsManager("config.json");
    metricsManager.RegisterSummary("Aboba", {0.3, 0.5, 0.95, 0.99});
    metricsManager.GetSummary("Aboba")->Observe(123);
    metricsManager.GetSummary("Aboba")->Observe(1);
    metricsManager.GetSummary("Aboba")->Observe(5);
    return 0;
}
