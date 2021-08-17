package models

type RequestModel struct {
	Users                int64       `json:"users"`
	ExecutionTimeInSeconds int64       `json:"executionTimeInSeconds"`
	RampUpTimeInSeconds    int64       `json:"rampUpTimeInSeconds"`
	HTTPRequest            HTTPRequest `json:"httpRequest"`
}

type HTTPRequest struct {
	Method             string            `json:"method"`
	URL    string            `json:"url"`
	Body            interface{}       `json:"body"`
	Headers            map[string]string `json:"headers"`
	QueryParams        map[string]string `json:"queryParams"`
}

