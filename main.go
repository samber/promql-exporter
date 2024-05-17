package main

import (
	"bytes"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/alecthomas/kingpin/v2"
	"github.com/samber/promql_exporter/exporter"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"

	endpoint      = kingpin.Flag("endpoint", "PromQL http endpoint").Envar("ENDPOINT").Required().String()
	headers       = kingpin.Flag("header", "PromQL http header").Envar("HEADER").Strings()
	namespace     = kingpin.Flag("namespace", "Namespace for metrics").Envar("PROMQL_EXPORTER_NAMESPACE").Default("promql").String()
	listenAddress = kingpin.Flag("web.listen-address", "Address to listen on for web interface and telemetry.").Envar("PROMQL_EXPORTER_WEB_LISTEN_ADDRESS").Default(":9517").String()
	metricPath    = kingpin.Flag("web.telemetry-path", "Path under which to expose metrics.").Envar("PROMQL_EXPORTER_WEB_TELEMETRY_PATH").Default("/metrics").String()
	logFormat     = kingpin.Flag("log.format", "Log format, valid options are txt and json").Envar("PROMQL_EXPORTER_LOG_FORMAT").Default("txt").String()
	// logLevel       = kingpin.Flag("log.level", "Log level").Envar("PROMQL_EXPORTER_LOG_FORMAT").Default("debug").String()
)

func main() {
	kingpin.Version(version)
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	switch *logFormat {
	case "json":
		slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, nil)))
	default:
		slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, nil)))
	}

	headersKV := map[string]string{}
	for _, header := range *headers {
		parts := strings.SplitN(header, ":", 2)
		if len(parts) != 2 {
			slog.Warn(fmt.Sprintf("Invalid header: %s", header))
			continue
		}

		headersKV[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
	}

	slog.Info("PromQL Metrics Exporter",
		slog.Any("build.time", date),
		slog.Any("build.release", version),
		slog.Any("build.commit", commit),
		slog.Any("go.version", runtime.Version()),
		slog.Any("go.os", runtime.GOOS),
		slog.Any("go.arch", runtime.GOARCH))

	slog.Info(fmt.Sprintf("Providing metrics at %s%s", *listenAddress, *metricPath))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//nolint:errcheck
		w.Write([]byte(`<html>
			<head><title>PromQL Exporter</title></head>
			<body>
			<h1>PromQL Exporter</h1>
			<p><a href="` + *metricPath + `">Metrics</a></p>
			</body>
			</html>`))
	})

	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		result, err := exporter.GetMetrics(*endpoint, headersKV)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("error: %s", err)))
			return
		}

		count := 0
		buf := bytes.NewBuffer([]byte{})

		for i := range result.Data.Result {
			name, ok := result.Data.Result[i].Metric["__name__"]
			if !ok {
				continue
			}

			delete(result.Data.Result[i].Metric, "__name__")

			labels := []string{}
			for k, v := range result.Data.Result[i].Metric {
				labels = append(labels, fmt.Sprintf("%s=\"%s\"", k, v))
			}

			value := result.Data.Result[i].Value[1].(string)

			buf.WriteString(fmt.Sprintf("%s_%s{%s} %s\n", *namespace, name, strings.Join(labels, ", "), value))
			count++
		}

		slog.Info(fmt.Sprintf("Extracted %d metrics", count))
		w.WriteHeader(http.StatusOK)
		w.Write(buf.Bytes())
	})

	log.Fatal(http.ListenAndServe(*listenAddress, nil))
}
