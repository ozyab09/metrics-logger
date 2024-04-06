#include <iostream>
#include "json.hpp"

#include <map>
#include <any>
#include <string>

#include "logger_base.h"
#include "metrics.h"
#include <iostream>


int main() {

    MetricsManager metricsManager("/Users/abdullinsaid/CLionProjects/metrics-logger/Cpp_logger/Cpp_logger/config.json");
    MetricsLogger::BaseLogger logger;
    logger.Log(MetricsLogger::LogLevel::ERROR, "ABOBA");
    logger.SetLogLevel(MetricsLogger::LogLevel::FATAL);
    logger.Log(MetricsLogger::LogLevel::ERROR, "ABOB2");

    return 0;
}
