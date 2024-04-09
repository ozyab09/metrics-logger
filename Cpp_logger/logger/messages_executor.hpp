#pragma once

#include <functional>
#include <future>
#include <memory>
#include <thread>

#include "../concurrentqueue.h"

namespace MetricsLogger {
    class MessagesExecutor {
    public:
        MessagesExecutor() : ExecutorThread(ExecutorThreadFunction) {}
        ~MessagesExecutor() {
            AddMessageSending([=]() { IsRunning = false; });
            ExecutorThread.join();
        }
        void AddMessageSending(std::function<void()> SendMessage) {
            MessagesSendingQueue.enqueue(SendMessage);
        }
        size_t size_approx() {
            return MessagesSendingQueue.size_approx();
        }

    private:
        bool IsRunning{true};
        std::function<void()> ExecutorThreadFunction{[=]() {
            while (IsRunning) {
                std::function<void()> SendMessage;
                if (MessagesSendingQueue.try_dequeue(SendMessage)) {
                    SendMessage();
                } else {
                    std::this_thread::sleep_for(std::chrono::milliseconds(5));
                }
            }
        }};

    private:
        std::thread ExecutorThread;
        moodycamel::ConcurrentQueue<std::function<void()>> MessagesSendingQueue;
    };

} // namespace MetricsLogger