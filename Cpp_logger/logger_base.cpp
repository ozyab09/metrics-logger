#include "logger_base.h"

#include <sstream>
#include <iostream>

namespace MetricsLogger {
    std::string LogLevelToString(LogLevel level) {
        std::vector<std::string> levels = {"TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL", "NO"};
        return levels[level];
    }

    void BaseLogger::Log(const LogLevel& logLevel, const std::string& message, const std::map<std::string, std::any>& otherFields) {
        if (logLevel > Level) {
            Executor.AddMessageSending([=]() {
                LogMessage logMessage;
                logMessage.Fields["Message"] = message;
                std::time_t now = std::chrono::system_clock::to_time_t(std::chrono::system_clock::now());
                std::string timestamp(30, '\0');
                std::strftime(&timestamp[0], timestamp.size(), "%Y-%m-%d %H:%M:%S", std::localtime(&now));
                logMessage.Fields["Timestamp"] = timestamp;
                logMessage.Fields["LogLevel"] = LogLevelToString(logLevel);
                std::ostringstream ss;
                ss << std::this_thread::get_id();
                logMessage.Fields["TreadId"] = ss.str();

                std::cout << std::any_cast<std::string>(logMessage.Fields["Timestamp"]) << ' '
                          << std::any_cast<std::string>(logMessage.Fields["LogLevel"]) << ' '
                          << std::any_cast<std::string>(logMessage.Fields["Message"]) << ' '
                          << std::any_cast<std::string>(logMessage.Fields["TreadId"]);
            });
        }
    }
}