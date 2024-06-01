#pragma once

#include "logger_base.h"
#include <functional>
#include <utility>

namespace MetricsLogger {
    class CustomConsoleLogger : BaseLogger {
    public:
        explicit CustomConsoleLogger(std::function<void(const LogLevel& logLevel,
                                                        const std::string& message,
                                                        const std::map<std::string, std::any>& otherFields)> func) : ConsoleLoggingFunction(std::move(func)) {}

        void Log(
                const LogLevel& logLevel,
                const std::string& message,
                const std::map<std::string, std::any>& otherFields = {}
        ) override;

        void SetConsoleLoggingFunction(std::function<void(const LogLevel& logLevel,
                                                          const std::string& message,
                                                          const std::map<std::string, std::any>& otherFields)> func);

    private:
        std::function<void(const LogLevel& logLevel,
                           const std::string& message,
                           const std::map<std::string, std::any>& otherFields)> ConsoleLoggingFunction;
    };
}
