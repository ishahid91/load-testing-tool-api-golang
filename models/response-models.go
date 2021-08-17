package models

type CurrentStats struct {
	ResponseTime int64
	Error        bool
}

type LoadTestStats struct {
	AverageResponseTime int64 `json:"average_response_time"`
	MaxResponseTime    int64 `json:"peak_response_time"`
	TotalRequests       int64 `json:"total_requests"`
	ErrorCount          int64 `json:"error_count"`
}