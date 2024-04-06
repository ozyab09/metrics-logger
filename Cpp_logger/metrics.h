#include <cpprest/http_listener.h>
#include <cpprest/json.h>
#include <cpprest/uri.h>
#include <cpprest/http_client.h>
#include <cpprest/producerconsumerstream.h>


using namespace web;
using namespace web::http;
using namespace web::http::experimental::listener;

class MetricsManager {
public:
    explicit MetricsManager(const std::string& path);

    void SetMetric(const std::string& name, const std::string& value);

    void SetMetric(const std::string& name, double_t value);

    void SetMetric(const std::string& name, const json::value& value);

    json::value GetMetric(const std::string& name);

    json::value& MutableMetric(const std::string& name);

    json::value GetMetricsJson();
private:
    void HandleGet(const http_request& request);
private:
    json::value Config;
    json::value Metrics;
    http_listener Listener;
};
