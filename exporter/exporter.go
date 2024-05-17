package exporter

type Metrics struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric map[string]string `json:"metric"`
			Value  []any             `json:"value"`
		} `json:"result"`
	} `json:"data"`
}

func GetMetrics(endpoint string, headers map[string]string) (*Metrics, error) {
	return HTTPRequest[*Metrics](endpoint, `/api/v1/query?query={__name__!=""}`, headers)
}
