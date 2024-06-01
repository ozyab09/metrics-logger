#include "custom_console_logger.h"


namespace MetricsLogger {
        void CustomConsoleLogger::Log(
                const LogLevel& logLevel,
                const std::string& message,
                const std::map<std::string, std::any>& otherFields
        ) {
            if (logLevel >= Level) {
                Executor.AddMessageSending([=]() {
                    ConsoleLoggingFunction(logLevel, message, otherFields);
                });
            }
        }

        void CustomConsoleLogger::SetConsoleLoggingFunction(std::function<void(const LogLevel& logLevel,
                                                          const std::string& message,
                                                          const std::map<std::string, std::any>& otherFields)> func) {
            ConsoleLoggingFunction = std::move(func);
        }
}
