#include "metrics.h"

namespace MetricsManager {
    void MetricsManager::HandleGet(const http_request &request) {
        std::cout << "ABOBA" << std::endl;
        request.reply(status_codes::OK, SerializeMetrics());
    }

    MetricsManager::~MetricsManager() {
        Listener.close().wait();
    }

    MetricsManager::MetricsManager(const std::string &path) {
        std::ifstream file(path);

        try {
            Config = web::json::value::parse(file);
        } catch (const web::json::json_exception &e) {
            std::cerr << "Metrics' config parsing error: " << e.what() << std::endl;
            throw;
        }

        file.close();

        Listener = http_listener(
                "http://localhost:" + std::to_string(Config["port"].as_integer()) + Config["metrics_path"].as_string());
        std::cout << "http://localhost:" + std::to_string(Config["port"].as_integer()) + Config["metrics_path"].as_string() << std::endl;
        Listener.support(methods::GET, [this](const http_request &http_req) { this->HandleGet(http_req); });
        Listener.open().wait();
    }

    std::string MetricsManager::SerializeMetrics() {
        std::unique_lock<std::mutex> lock(Mutex);
        if (Config["output_format"].as_string() == "prometheus") {
            std::string result;
            for (const auto& it : Counters) {
                result += it.second->SerializeInPrometheus();
            }
            for (const auto& it : Gauges) {
                result += it.second->SerializeInPrometheus();
            }
            for (const auto& it : Summaries) {
                result += it.second->SerializeInPrometheus();
            }
            return result;
        } else if (Config["output_format"].as_string() == "json") {
            web::json::value result;
            for (const auto& it : Counters) {
                for (const auto& field : it.second->SerializeInJson().as_object()) {
                    result[field.first] = field.second;
                }
            }
            for (const auto& it : Gauges) {
                for (const auto& field : it.second->SerializeInJson().as_object()) {
                    result[field.first] = field.second;
                }
            }
            for (const auto& it : Summaries) {
                for (const auto& field : it.second->SerializeInJson().as_object()) {
                    result[field.first] = field.second;
                }
            }
            return result.serialize();
        }
        std::cerr << "Bad output_format type" << std::endl;
        return "";
    }

    std::shared_ptr<Counter> MetricsManager::GetCounter(const std::string &name) {
        return Counters[name];
    }

    std::shared_ptr<Gauge> MetricsManager::GetGauge(const std::string &name) {
        return Gauges[name];
    }

    std::shared_ptr<Summary> MetricsManager::GetSummary(const std::string &name) {
        return Summaries[name];
    }

    void MetricsManager::RegisterCounter(const std::string& label) {
        if (CheckLabel(label)) {
            Counters[label] = std::make_shared<Counter>(label);
        }
    }

    void MetricsManager::RegisterGauge(const std::string& label) {
        if (CheckLabel(label)) {
            Gauges[label] = std::make_shared<Gauge>(label);
        }
    }

    void MetricsManager::RegisterSummary(const std::string& label, const std::vector<double>& percentiles, size_t cacheSize) {
        if (CheckLabel(label)) {
            Summaries[label] = std::make_shared<Summary>(label, percentiles, cacheSize);
        }
    }

    bool MetricsManager::CheckLabel(const std::string& label) {
        if (Counters.count(label) != 0 && Gauges.count(label) != 0 && Summaries.count(label) != 0) {
            std::cerr << "There is already metric with name " << label << std::endl;
            return false;
        }
        if (label.empty()) {
            std::cerr << "Empty label" << std::endl;
            return false;
        }
        if (!(label.front() >= 'A' && label.front() <= 'Z' ||
              label.front() >= 'a' && label.front() <= 'z')) {
            std::cerr << "Label must begin with a latin letter" << std::endl;
            return false;
        }
        for (const auto& c : label) {
            if (!(c >= 'A' && c <= 'Z' ||
                  c >= 'a' && c <= 'z' ||
                  c >= '0' && c <= '9' ||
                  c == '_')) {
                std::cerr << "Incorrect char in label: " << c << std::endl;
                return false;
            }
        }
        return true;
    }

}