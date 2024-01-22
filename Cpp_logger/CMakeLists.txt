cmake_minimum_required(VERSION 3.25)
project(Cpp_logger)

set(CMAKE_CXX_STANDARD 17)

add_executable(Cpp_logger main.cpp logger_base.cpp logger_base.h concurrentqueue.h messages_executor.hpp log_message.h log_level.h)
