<p align="center">
  <img src="https://raw.githubusercontent.com/PKief/vscode-material-icon-theme/ec559a9f6bfd399b82bb44393651661b08aaf7ba/icons/folder-markdown-open.svg" width="100" alt="project-logo">
</p>
<p align="center">
    <h1 align="center">Metrics-logger</h1>
</p>
<p align="center">
    <em><code> Logging library that generates application metrics</code></em>
</p>

<br><!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary><br>

- [ Overview](#-overview)
- [ Features](#-features)
- [ Repository Structure](#-repository-structure)
- [ Modules](#-modules)
- [ System Requirements](#-system-requirements)
- [ Installation](#-installation)
- [ Usage](#-usage)
</details>
<hr>

##  Overview

<code>► This product is designed for software developers who want a simple and convenient way to create logs of their projects. It can also be useful for those who need to collect and analyze metrics about their software.</code>

---

##  Features

<code>► We are planning to add new features to our library in order to make it more user-friendly and convenient. We would appreciate any suggestions from the open-source community that could help us improve our product.</code>

---

##  Repository Structure

```sh
└── metrics-logger/
    ├── Cpp_logger
    │   ├── CMakeLists.txt
    │   ├── README.md
    │   ├── concurrentqueue.h
    │   ├── config.json
    │   ├── json.hpp
    │   ├── logger
    │   ├── main.cpp
    │   └── metrics
    ├── README.md
    └── go_metrics_logger
        ├── README.md
        ├── go.mod
        ├── go.sum
        ├── logger
        ├── metrics
        └── metrics_logger.go
```

---

##  Modules

<details closed><summary>Cpp_logger</summary>

| File                                                                                                    | Summary                         |
| ---                                                                                                     | ---                             |
| [json.hpp](https://github.com/ozyab09/metrics-logger/blob/master/Cpp_logger/json.hpp)                   | <code>► INSERT-TEXT-HERE</code> |
| [config.json](https://github.com/ozyab09/metrics-logger/blob/master/Cpp_logger/config.json)             | <code>► INSERT-TEXT-HERE</code> |
| [CMakeLists.txt](https://github.com/ozyab09/metrics-logger/blob/master/Cpp_logger/CMakeLists.txt)       | <code>► INSERT-TEXT-HERE</code> |
| [concurrentqueue.h](https://github.com/ozyab09/metrics-logger/blob/master/Cpp_logger/concurrentqueue.h) | <code>► INSERT-TEXT-HERE</code> |
| [main.cpp](https://github.com/ozyab09/metrics-logger/blob/master/Cpp_logger/main.cpp)                   | <code>► INSERT-TEXT-HERE</code> |

</details>

<details closed><summary>Cpp_logger.logger</summary>

| File                                                                                                                   | Summary                         |
| ---                                                                                                                    | ---                             |
| [logger_base.h](https://github.com/ozyab09/metrics-logger/blob/master/Cpp_logger/logger/logger_base.h)                 | <code>► INSERT-TEXT-HERE</code> |
| [log_message.h](https://github.com/ozyab09/metrics-logger/blob/master/Cpp_logger/logger/log_message.h)                 | <code>► INSERT-TEXT-HERE</code> |
| [log_level.h](https://github.com/ozyab09/metrics-logger/blob/master/Cpp_logger/logger/log_level.h)                     | <code>► INSERT-TEXT-HERE</code> |
| [logger_base.cpp](https://github.com/ozyab09/metrics-logger/blob/master/Cpp_logger/logger/logger_base.cpp)             | <code>► INSERT-TEXT-HERE</code> |
| [messages_executor.hpp](https://github.com/ozyab09/metrics-logger/blob/master/Cpp_logger/logger/messages_executor.hpp) | <code>► INSERT-TEXT-HERE</code> |

</details>

<details closed><summary>Cpp_logger.metrics</summary>

| File                                                                                                    | Summary                         |
| ---                                                                                                     | ---                             |
| [counter.h](https://github.com/ozyab09/metrics-logger/blob/master/Cpp_logger/metrics/counter.h)         | <code>► INSERT-TEXT-HERE</code> |
| [gauge.h](https://github.com/ozyab09/metrics-logger/blob/master/Cpp_logger/metrics/gauge.h)             | <code>► INSERT-TEXT-HERE</code> |
| [metrics.h](https://github.com/ozyab09/metrics-logger/blob/master/Cpp_logger/metrics/metrics.h)         | <code>► INSERT-TEXT-HERE</code> |
| [summary.h](https://github.com/ozyab09/metrics-logger/blob/master/Cpp_logger/metrics/summary.h)         | <code>► INSERT-TEXT-HERE</code> |
| [metrics.cpp](https://github.com/ozyab09/metrics-logger/blob/master/Cpp_logger/metrics/metrics.cpp)     | <code>► INSERT-TEXT-HERE</code> |
| [counter.cpp](https://github.com/ozyab09/metrics-logger/blob/master/Cpp_logger/metrics/counter.cpp)     | <code>► INSERT-TEXT-HERE</code> |
| [summary.cpp](https://github.com/ozyab09/metrics-logger/blob/master/Cpp_logger/metrics/summary.cpp)     | <code>► INSERT-TEXT-HERE</code> |
| [gauge.cpp](https://github.com/ozyab09/metrics-logger/blob/master/Cpp_logger/metrics/gauge.cpp)         | <code>► INSERT-TEXT-HERE</code> |
| [base_metric.h](https://github.com/ozyab09/metrics-logger/blob/master/Cpp_logger/metrics/base_metric.h) | <code>► INSERT-TEXT-HERE</code> |

</details>

<details closed><summary>go_metrics_logger</summary>

| File                                                                                                           | Summary                         |
| ---                                                                                                            | ---                             |
| [go.sum](https://github.com/ozyab09/metrics-logger/blob/master/go_metrics_logger/go.sum)                       | <code>► INSERT-TEXT-HERE</code> |
| [go.mod](https://github.com/ozyab09/metrics-logger/blob/master/go_metrics_logger/go.mod)                       | <code>► INSERT-TEXT-HERE</code> |
| [metrics_logger.go](https://github.com/ozyab09/metrics-logger/blob/master/go_metrics_logger/metrics_logger.go) | <code>► INSERT-TEXT-HERE</code> |

</details>

<details closed><summary>go_metrics_logger.logger</summary>

| File                                                                                                                        | Summary                         |
| ---                                                                                                                         | ---                             |
| [messages_executor.go](https://github.com/ozyab09/metrics-logger/blob/master/go_metrics_logger/logger/messages_executor.go) | <code>► INSERT-TEXT-HERE</code> |
| [internal_logger.go](https://github.com/ozyab09/metrics-logger/blob/master/go_metrics_logger/logger/internal_logger.go)     | <code>► INSERT-TEXT-HERE</code> |
| [message.go](https://github.com/ozyab09/metrics-logger/blob/master/go_metrics_logger/logger/message.go)                     | <code>► INSERT-TEXT-HERE</code> |
| [file_logger.go](https://github.com/ozyab09/metrics-logger/blob/master/go_metrics_logger/logger/file_logger.go)             | <code>► INSERT-TEXT-HERE</code> |

</details>

<details closed><summary>go_metrics_logger.metrics</summary>

| File                                                                                                                         | Summary                         |
| ---                                                                                                                          | ---                             |
| [metrics.go](https://github.com/ozyab09/metrics-logger/blob/master/go_metrics_logger/metrics/metrics.go)                     | <code>► INSERT-TEXT-HERE</code> |
| [metrics_publisher.go](https://github.com/ozyab09/metrics-logger/blob/master/go_metrics_logger/metrics/metrics_publisher.go) | <code>► INSERT-TEXT-HERE</code> |

</details>

---

##  System Requirements

* **C++**: `version x.y.z`
* **GoLang**: `version x.y.z`
* **Java**: `version x.y.z`
---

##  Installation



<h3>From <code>github</code></h3>

<h4> For GoLang: </h4>

> ```
> $ go get https://github.com/ozyab09/metrics-logger/go_metrics_logger@latest
> ```

<h4> For C++: </h4>

> ```
> $ c++ get https://github.com/ozyab09/metrics-logger/cpp_logger@latest
> ```

<h4> For Java: </h4>

> ```
> $ java ...
> ```

###  Usage

<code>The logging library "Metrics-logger" is designed for use in software development. It helps to monitor services and reduces the amount of code that programmers need to write. The target audience for this software product includes information system developers of all levels, who are involved in writing and maintaining services. </code>


