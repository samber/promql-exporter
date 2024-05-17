
# PromQL Exporter

[![tag](https://img.shields.io/github/tag/samber/promql-exporter.svg)](https://github.com/samber/promql-exporter/releases)
![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.18.0-%23007d9c)
[![GoDoc](https://godoc.org/github.com/samber/promql-exporter?status.svg)](https://pkg.go.dev/github.com/samber/promql-exporter)
![Build Status](https://github.com/samber/promql-exporter/actions/workflows/test.yml/badge.svg)
[![Go report](https://goreportcard.com/badge/github.com/samber/promql-exporter)](https://goreportcard.com/report/github.com/samber/promql-exporter)
[![Coverage](https://img.shields.io/codecov/c/github/samber/promql-exporter)](https://codecov.io/gh/samber/promql-exporter)
[![Contributors](https://img.shields.io/github/contributors/samber/promql-exporter)](https://github.com/samber/promql-exporter/graphs/contributors)
[![License](https://img.shields.io/github/license/samber/promql-exporter)](./LICENSE)

> A Prometheus Exporter for PromQL-compatible endpoints

Some cloud providers do not offer federation endpoint or remote write. This exporter aims to export metrics using the /query API endpoint.

## 🚀 Run

Using Docker:

```sh
docker run --rm -it -p 9517:9517 -e ENDPOINT=http://demo.robustperception.io:9090 samber/promql-exporter:v0.1.0
```

Or using a binary:

```sh
wget -O promql_exporter https://github.com/samber/promql-exporter/releases/download/v0.1.0/promql_exporter.1.1_linux_amd64
chmod +x promql_exporter
./promql_exporter --endpoint xxxx --header 'x-token: yyyy' --header 'x-token: zzzz'
```

## 💡 Usage

```sh
./promql_exporter
usage: promql_exporter --endpoint=http://demo.robustperception.io:9090 [<flags>]

Flags:
  -h, --help                           Show context-sensitive help (also try --help-long and --help-man).
      --endpoint                       PromQL http endpoint ($ENDPOINT)
      --header                         PromQL http header ($HEADER)
      --namespace="promql"             Namespace for metrics ($PROMQL_EXPORTER_NAMESPACE)
      --web.listen-address=":9517"     Address to listen on for web interface and telemetry. ($PROMQL_EXPORTER_WEB_LISTEN_ADDRESS)
      --web.telemetry-path="/metrics"  Path under which to expose metrics. ($PROMQL_EXPORTER_WEB_TELEMETRY_PATH)
      --log.format="txt"               Log format, valid options are txt and json ($PROMQL_EXPORTER_LOG_FORMAT)
      --version                        Show application version.
```

## 🤝 Contributing

- Ping me on Twitter [@samuelberthe](https://twitter.com/samuelberthe) (DMs, mentions, whatever :))
- Fork the [project](https://github.com/samber/promql-exporter)
- Fix [open issues](https://github.com/samber/promql-exporter/issues) or request new features

Don't hesitate ;)

```bash
# Install some dev dependencies
make tools

# Run tests
make test
# or
make watch-test
```

## 👤 Contributors

![Contributors](https://contrib.rocks/image?repo=samber/promql-exporter)

## 💫 Show your support

Give a ⭐️ if this project helped you!

[![GitHub Sponsors](https://img.shields.io/github/sponsors/samber?style=for-the-badge)](https://github.com/sponsors/samber)

## 📝 License

Copyright © 2024 [Samuel Berthe](https://github.com/samber).

This project is [MIT](./LICENSE) licensed.
