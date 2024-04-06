cmake_minimum_required(VERSION 3.9)
project(Cpp_logger)

set(CMAKE_CXX_STANDARD 17)
find_package(cpprestsdk REQUIRED)
find_package(OpenSSL REQUIRED)

add_executable(Cpp_logger main.cpp logger_base.cpp logger_base.h concurrentqueue.h messages_executor.hpp log_message.h log_level.h metrics.cpp metrics.h)
target_link_libraries(Cpp_logger PRIVATE OpenSSL::SSL OpenSSL::Crypto PRIVATE cpprestsdk::cpprest)

