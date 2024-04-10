**Краткий обзор**

```
// Создание логгера
Logger logger = new Logger();
```

```
// Создание метрик
Metrics metrics = new Metrics();
```
```
// Открыть метрики и логи на порт
Config config = new Config("SOME_SERVICE", "logs", "metrics", 8080);
Reporter.start(config, logger, metrics);
```
```
// Пример использования логгера
logger.d("log1");
logger.d("log2");
```

Логирование доступно с уровнями\
`DEBUG, ERROR, VERBOSE, INFO, WARNING`
```
// Примеры создания и редактирования метрик

// Gauge
metrics.createCounter("counter_ex");
metrics.getCounter("counter_ex").increment();

// Gauge
metrics.createGauge("gauge_ex");
metrics.getGauge("gauge_ex").set(45);

// Summary

// name: String, capacity: Int, percentiles: Double...
metrics.createSummary("summary_ex", 10, 50);
Summary summary = metrics.getSummary("summary_ex")
summary.add(1);
summary.add(5);
summary.add(-12);
summary.add(100);
```
