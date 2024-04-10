#pragma once
#include <cpprest/http_listener.h>
#include <cpprest/json.h>
#include <cpprest/uri.h>
#include <cpprest/http_client.h>
#include <cpprest/producerconsumerstream.h>

#include "counter.h"
#include "gauge.h"
#include "summary.h"


using namespace web;
using namespace web::http;
using namespace web::http::experimental::listener;


namespace MetricsManager {
    class MetricsManager {
    public:
        explicit MetricsManager(const std::string& path);

        ~MetricsManager();

        void RegisterCounter(const std::string& label);

        void RegisterGauge(const std::string& label);

        void RegisterSummary(const std::string& label, const std::vector<double>& percentiles, size_t cacheSize = 1024);

        std::string SerializeMetrics();

        std::shared_ptr<Counter> GetCounter(const std::string &label);

        std::shared_ptr<Gauge> GetGauge(const std::string &label);

        std::shared_ptr<Summary> GetSummary(const std::string &label);
    private:
        void HandleGet(const http_request& request);

        bool CheckLabel(const std::string& label);
    private:
        json::value Config;
        http_listener Listener;
        std::unordered_map<std::string, std::shared_ptr<Counter>> Counters;
        std::unordered_map<std::string, std::shared_ptr<Gauge>> Gauges;
        std::unordered_map<std::string, std::shared_ptr<Summary>> Summaries;
        mutable std::mutex Mutex;
    };
}