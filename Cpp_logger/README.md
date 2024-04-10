# C++ logger and metrics library

## Logger

Thread-safe cross-platform logger library with added metrics, such as timestamp, loglevel etc.

BaseLogger class can be inherited to use custom log functions.

### Log levels

TRACE, DEBUG, INFO, WARN, ERROR, FATAL, NO.

### Usage examples

#### Log in console
```cpp
MetricsLogger::BaseLogger logger;
logger.Log(MetricsLogger::LogLevel::ERROR, "Log in console");
// Result in console:
// 2024-03-10 20:17:14 ERROR Log in console 0x16d69b000
```

#### Log in console in json format
```cpp
MetricsLogger::BaseLogger logger;
logger.LogJson(MetricsLogger::LogLevel::ERROR, "Log in console");
// Result in console:
// {"LogLevel":"ERROR","Message":"Log in console","ThreadId":"","Timestamp":"2024-03-10 21:23:08","TreadId":"0x16f3d3000"}
```

#### Log in file in json format
```cpp
MetricsLogger::BaseLogger logger;
std::any val = std::string("log value");
logger.LogJsonInFile(MetricsLogger::LogLevel::ERROR, "Log in console", "file.txt", {{"Optional log", val}});
// Result in file:
// {"LogLevel":"ERROR","Message":"Log in console","Optional log":"log value","ThreadId":"0x16b727000","Timestamp":"2024-03-10 21:55:56"}
```

#### Log in file
```cpp
MetricsLogger::BaseLogger logger;
std::any val = std::string("log value");
logger.LogInFile(MetricsLogger::LogLevel::ERROR, "Log in console", "file.txt", {{"Optional log", val}});

// Result in file:
// 2024-03-10 21:58:08 ERROR Log in console 0x16b8a3000 log value
```

## Metrics

Thread-safe cross-platform metrics library
3 Metric types: 
* Counter
* Gauge
* Summary

### Structure

MetricsManager - base class for counters registration, needs json config

### Config fields

"port" - port for metrics

"metrics_path" - path for metrics on port

"output_format": serialization format: "prometheus" or "json"

### Usage examples

#### Counter
```cpp
MetricsManager::MetricsManager metricsManager("config.json");
metricsManager.RegisterCounter("Aboba"); // "Aboba" value = 0
metricsManager.GetCounter("Aboba")->Increment(); // "Aboba" value = 1
metricsManager.GetCounter("Aboba")->Increment(5); // "Aboba" value = 6
metricsManager.GetCounter("Aboba")->Reset();  // "Aboba" value = 0
```

#### Gauge
```cpp
MetricsManager::MetricsManager metricsManager("config.json");
metricsManager.RegisterGauge("Aboba"); // "Aboba" value = 0
metricsManager.GetGauge("Aboba")->Increment(); // "Aboba" value = 1
metricsManager.GetGauge("Aboba")->Decrement(23); // "Aboba" value = -22
metricsManager.GetGauge("Aboba")->Set(0); // "Aboba" value = 0
```

#### Summary
```cpp
MetricsManager::MetricsManager metricsManager("config.json");
metricsManager.RegisterSummary("Aboba", {0.3, 0.5, 0.95, 0.99}); // "Aboba" summary would calc {0.3, 0.5, 0.95, 0.99} percentiles
```



  
