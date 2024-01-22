#pragma once

#include "log_level.h"
#include "log_message.h"
#include "messages_executor.hpp"

namespace MetricsLogger {
    class BaseLogger {
    public:
        void Log(
            const LogLevel& logLevel,
            const std::string& message,
            const std::map<std::string, std::any>& otherFields = {}
        );

    private:
        MessagesExecutor Executor;
        std::atomic<LogLevel> Level{LogLevel::WARN};
    };
}
