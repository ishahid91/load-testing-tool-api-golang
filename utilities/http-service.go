package utilities

import (
	"github.com/tcnksm/go-httpstat"
	"io"
	"log"
	"net/http"
	"time"
)
type HTTPClient struct {
	httpClient        *http.Client
	HTTPClientService HTTPClientService
}

type HTTPClientService interface {
	Initialise()
	SendRequest(string, string, map[string]string, map[string]string, io.Reader) (*http.Response, error)
}

func (i *HTTPClient) Initialise() {
	i.httpClient = &http.Client{}
}

func (i *HTTPClient) SendRequest(method string, url string, queryParams map[string]string, headers map[string]string, body io.Reader) (resp *http.Response, responseTimeInMilliSeconds int64, err error) {

	start := time.Now()
	var resultMetrics httpstat.Result
	req, err := http.NewRequest(method,url,body)
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


