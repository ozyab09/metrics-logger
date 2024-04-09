#pragma once

#include <string>

namespace MetricsLogger {
    enum LogLevel {
        TRACE = 0,
        DEBUG,
        INFO,
        WARN,
        ERROR,
        FATAL,
        NO
    };
}
