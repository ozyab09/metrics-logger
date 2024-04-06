#include "metrics.h"

void MetricsManager::HandleGet(const http_request& request) {
    request.reply(status_codes::OK, Metrics);
}

MetricsManager::MetricsManager(const std::string& path) {
    std::ifstream file(path);

    try {
        Config = web::json::value::parse(file);
    } catch (const web::json::json_exception& e) {
        std::cerr << "Metrics config parsing error: " << e.what() << std::endl;
    }

    file.close();

    http_listener listener("http://localhost:" + std::to_string(Config["port"].as_integer()) + "/metrics");
    listener.support(methods::GET, [this](const http_request& http_req) { this->HandleGet(http_req); });
}

void MetricsManager::SetMetric(const std::string& name, const std::string& value) {
    Metrics[name] = json::value::string(value);
}

void MetricsManager::SetMetric(const std::string& name, double_t value) {
    Metrics[name] = json::value::number(value);
}

void MetricsManager::SetMetric(const std::string& name, const json::value& value) {
    Metrics[name] = value;
}

json::value MetricsManager::GetMetric(const std::string& name) {
    return Metrics[name];
}

json::value& MetricsManager::MutableMetric(const std::string& name) {
    return Metrics[name];
}

json::value MetricsManager::GetMetricsJson() {
    return Metrics;
}
