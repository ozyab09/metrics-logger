#pragma once

#include "log_level.h"
#include "log_message.h"
#include "messages_executor.hpp"

namespace MetricsLogger {
    class BaseLogger {
    public:
        virtual void Log(
            const LogLevel& logLevel,
            const std::string& message,
            const std::map<std::string, std::any>& otherFields = {}
        );

        void LogJson(
                const LogLevel& logLevel,
                const std::string& message,
                const std::map<std::string, std::any>& otherFields = {}
        );

        void LogJsonInFile(
                const LogLevel& logLevel,
                const std::string& message,
                const std::string& filePath,
                const std::map<std::string, std::any>& otherFields = {}
        );

        void LogInFile(
                const LogLevel& logLevel,
                const std::string& message,
                const std::string& filePath,
                const std::map<std::string, std::any>& otherFields = {}
        );

        void SetLogLevel(const LogLevel& logLevel);

    protected:
        static std::string GetDefaultLogString(const LogLevel& logLevel,
                                        const std::string& message,
                                        const std::map<std::string, std::any>& otherFields);

    protected:
        MessagesExecutor Executor;
        std::atomic<LogLevel> Level{LogLevel::WARN};
    };
}
