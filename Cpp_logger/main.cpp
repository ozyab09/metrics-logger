#include <iostream>
#include "json.hpp"

#include <map>
#include <any>
#include <string>

#include "logger_base.h"
#include <iostream>
#include <cpprest/http_listener.h>
#include <cpprest/json.h>
#include <cpprest/uri.h>
#include <cpprest/http_client.h>
#include <cpprest/producerconsumerstream.h>


int main() {
    MetricsLogger::BaseLogger logger;
    logger.Log(MetricsLogger::LogLevel::ERROR, "ABOBA");
    logger.SetLogLevel(MetricsLogger::LogLevel::FATAL);
    logger.Log(MetricsLogger::LogLevel::ERROR, "ABOB2");

    return 0;
}
