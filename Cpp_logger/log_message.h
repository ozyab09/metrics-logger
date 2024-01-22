#pragma once

#include <string>
#include <any>
#include <map>

namespace MetricsLogger {
    class LogMessage {
    public:
        LogMessage() {
            Fields["Message"] = "";
            Fields["Timestamp"] = "";
            Fields["ProcessName"] = "";
            Fields["LogLevel"] = "NO";
            Fields["ThreadId"] = "";
        }

        std::map<std::string, std::any> Fields;
    };

}