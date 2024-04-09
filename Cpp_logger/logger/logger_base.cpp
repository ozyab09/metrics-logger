#include "logger_base.h"

#include <sstream>
#include <iostream>
#include <fstream>

namespace MetricsLogger {
    std::string LogLevelToString(LogLevel level) {
        std::vector<std::string> levels = {"TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL", "NO"};
        return levels[level];
    }

    std::string BaseLogger::GetDefaultLogString(const LogLevel& logLevel,
                                                const std::string& message,
                                                const std::map<std::string, std::any>& otherFields) {
        LogMessage logMessage;
        logMessage.MainFields["Message"] = message;
        std::time_t now = std::chrono::system_clock::to_time_t(std::chrono::system_clock::now());
        std::string timestamp(30, '\0');
        std::strftime(&timestamp[0], timestamp.size(), "%Y-%m-%d %H:%M:%S", std::localtime(&now));
        logMessage.MainFields["Timestamp"] = timestamp;
        logMessage.MainFields["LogLevel"] = LogLevelToString(logLevel);
        std::ostringstream ss;
        ss << std::this_thread::get_id();
        logMessage.MainFields["ThreadId"] = ss.str();

        for (const auto& field : otherFields) {
            if (field.second.type() == typeid(int)) {
                logMessage.AdditionalFields[field.first] = std::any_cast<int>(field.second);
            } else if (field.second.type() == typeid(std::string)) {
                logMessage.AdditionalFields[field.first] = std::any_cast<std::string>(field.second);
            } else if (field.second.type() == typeid(bool)) {
                logMessage.AdditionalFields[field.first] = std::any_cast<bool>(field.second);
            } else if (field.second.type() == typeid(double)) {
                logMessage.AdditionalFields[field.first] = std::any_cast<double>(field.second);
            } else if (field.second.type() == typeid(std::vector<int>)) {
                logMessage.AdditionalFields[field.first] = std::any_cast<std::vector<int>>(field.second);
            } else if (field.second.type() == typeid(std::vector<double>)) {
                logMessage.AdditionalFields[field.first] = std::any_cast<std::vector<double>>(field.second);
            } else if (field.second.type() == typeid(std::vector<bool>)) {
                logMessage.AdditionalFields[field.first] = std::any_cast<std::vector<bool>>(field.second);
            }
        }

        std::string result;
        result += logMessage.MainFields["Timestamp"].get<const std::string>() + " "
                  + logMessage.MainFields["LogLevel"].get<const std::string>() + " "
                  + logMessage.MainFields["Message"].get<const std::string>() + " "
                  + logMessage.MainFields["ThreadId"].get<const std::string>();

        for (auto it = logMessage.AdditionalFields.begin(); it != logMessage.AdditionalFields.end(); ++it) {
            if (it.value().is_string()) {
                result += " " + it.value().get<const std::string>();
            } else {
                result += " " + it.value().dump();
            }
        }

        return result;
    }

    void BaseLogger::SetLogLevel(const LogLevel& logLevel) {
        Level = logLevel;
    }

    void BaseLogger::Log(const LogLevel& logLevel, const std::string& message, const std::map<std::string, std::any>& otherFields) {
        if (logLevel > Level) {
            Executor.AddMessageSending([=]() {
                std::cout << GetDefaultLogString(logLevel, message, otherFields) << std::endl;
            });
        }
    }

    void BaseLogger::LogJson(const LogLevel& logLevel, const std::string& message, const std::map<std::string, std::any>& otherFields) {
        if (logLevel > Level) {
            Executor.AddMessageSending([=]() {
                LogMessage logMessage;
                logMessage.MainFields["Message"] = message;
                std::time_t now = std::chrono::system_clock::to_time_t(std::chrono::system_clock::now());
                std::string timestamp(30, '\0');
                std::strftime(&timestamp[0], timestamp.size(), "%Y-%m-%d %H:%M:%S", std::localtime(&now));
                logMessage.MainFields["Timestamp"] = timestamp;
                logMessage.MainFields["LogLevel"] = LogLevelToString(logLevel);
                std::ostringstream ss;
                ss << std::this_thread::get_id();
                logMessage.MainFields["TreadId"] = ss.str();

                for (const auto& field : otherFields) {
                    if (field.second.type() == typeid(int)) {
                        logMessage.MainFields[field.first] = std::any_cast<int>(field.second);
                    } else if (field.second.type() == typeid(std::string)) {
                        logMessage.MainFields[field.first] = std::any_cast<std::string>(field.second);
                    } else if (field.second.type() == typeid(bool)) {
                        logMessage.MainFields[field.first] = std::any_cast<bool>(field.second);
                    } else if (field.second.type() == typeid(double)) {
                        logMessage.MainFields[field.first] = std::any_cast<double>(field.second);
                    } else if (field.second.type() == typeid(std::vector<int>)) {
                        logMessage.MainFields[field.first] = std::any_cast<std::vector<int>>(field.second);
                    } else if (field.second.type() == typeid(std::vector<double>)) {
                        logMessage.MainFields[field.first] = std::any_cast<std::vector<double>>(field.second);
                    } else if (field.second.type() == typeid(std::vector<bool>)) {
                        logMessage.MainFields[field.first] = std::any_cast<std::vector<bool>>(field.second);
                    }
                }

                std::cout << logMessage.MainFields << "/n";
            });
        }
    }

    void BaseLogger::LogInFile(const LogLevel& logLevel, const std::string& message, const std::string& filePath, const std::map<std::string, std::any>& otherFields) {
        if (logLevel > Level) {
            Executor.AddMessageSending([=]() {
                std::ofstream out;
                out.open(filePath);
                out << GetDefaultLogString(logLevel, message, otherFields) << std::endl;

                out.close();
            });
        }
    }
}