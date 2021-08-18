package utilities

import (
	"bytes"
	"encoding/json"
	"github.com/tcnksm/go-httpstat"
	"log"
	"net/http"
	"time"
)
type HTTPClient struct {
	httpClient        *http.Client
}

type IHTTPClientService interface {
	SendRequest(method string, url string, queryParams map[string]string, headers map[string]string, body interface{}) (resp *http.Response, responseTimeInMilliSeconds int64, err error)
}

func HTTPClientService() IHTTPClientService {
	return &HTTPClient{httpClient: &http.Client{}}
}

func (i *HTTPClient) SendRequest(method string, url string, queryParams map[string]string, headers map[string]string, body interface{}) (resp *http.Response, responseTimeInMilliSeconds int64, err error) {

	start := time.Now()
	var resultMetrics httpstat.Result
	payloadBytes, err := json.Marshal(body)
	if err != nil {
		log.Print("Unable to convert payload to bytes : ", err.Error())
	}
	payload := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest(method,url,payload)
	trace := httpstat.WithHTTPStat(req.Context(), &resultMetrics)
	req = req.WithContext(trace)
	if err != nil{
		log.Printf("Error in creating request : %s, %s", method,url)
		return
	}

	if len(headers) > 0 {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}

	if len(queryParams) > 0 {
		query := req.URL.Query()
		for key, value := range queryParams {
			query.Add(key, value)
		}
		req.URL.RawQuery = query.Encode()
	}

	resp, err = i.httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()
	resultMetrics.End(time.Now())

	responseTimeInMilliSeconds = time.Since(start).Milliseconds()
	return
}


