#include <iostream>
#include "json.hpp"

#include <map>
#include <any>
#include <string>

#include "logger/logger_base.h"
#include "metrics/metrics.h"
#include <iostream>


int main() {

    MetricsManager::MetricsManager metricsManager("/Users/abdullinsaid/CLionProjects/metrics-logger/Cpp_logger/Cpp_logger/config.json");
    metricsManager.RegisterCounter("Aboba");
    std::string s;
    getline(std::cin, s);
    metricsManager.GetCounter("Aboba")->Increment();
    getline(std::cin, s);
    return 0;
}
