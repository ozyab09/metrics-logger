#pragma once

#include <string>
#include <any>
#include <map>
#include "../json.hpp"

namespace MetricsLogger {
    class LogMessage {
    public:
        LogMessage() {
            MainFields["Message"] = "";
            MainFields["Timestamp"] = "";
            MainFields["LogLevel"] = "NO";
            MainFields["ThreadId"] = "";
        }

        nlohmann::json MainFields;
        nlohmann::json AdditionalFields;
    };

}