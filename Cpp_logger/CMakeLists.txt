cmake_minimum_required(VERSION 3.9)
project(Cpp_logger)

set(CMAKE_CXX_STANDARD 17)
find_package(cpprestsdk REQUIRED)
find_package(OpenSSL REQUIRED)

add_executable(Cpp_logger main.cpp logger/logger_base.cpp logger/logger_base.h concurrentqueue.h logger/messages_executor.hpp logger/log_message.h logger/log_level.h metrics/metrics.cpp metrics/metrics.h metrics/gauge.cpp metrics/gauge.h metrics/summary.cpp metrics/summary.h metrics/counter.cpp metrics/counter.h metrics/base_metric.h metrics/histogram.cpp metrics/histogram.h)
target_link_libraries(Cpp_logger PRIVATE OpenSSL::SSL OpenSSL::Crypto PRIVATE cpprestsdk::cpprest)

