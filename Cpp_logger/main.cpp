#include <iostream>
#include "json.hpp"

#include <map>
#include <any>
#include <string>

#include "logger_base.h"

int main() {
    MetricsLogger::BaseLogger logger;

    logger.Log(MetricsLogger::FATAL, "ABOBA");

    return 0;
}
