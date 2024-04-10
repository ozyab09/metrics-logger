package hse.serrriy;

public class Config {
    String applicationName;
    String logsPath;
    String metricsPath;
    int port;

    Config(
        String applicationName,
        String logsPath,
        String metricsPath,
        int port
    ) {
        this.applicationName = applicationName;
        this.logsPath = logsPath;
        this.metricsPath = metricsPath;
        this.port = port;
    }
}
