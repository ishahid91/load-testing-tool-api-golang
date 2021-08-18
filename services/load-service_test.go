package services

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"load-test-tool/models"
	"net/http"
	"testing"
)

type MockHttpClientService struct {
	mock.Mock
}

func (mock * MockHttpClientService) SendRequest(method string, url string, queryParams map[string]string,
	headers map[string]string, body interface{}) (resp *http.Response, responseTimeInMilliSeconds int64, err error) {

	args := mock.Called()
	return args.Get(0).(*http.Response), int64(args.Int(1)),args.Error(2)
}

func TestLoadTest(t *testing.T) {

	mockHTTPClient := new(MockHttpClientService)
	mockHTTPClient.On("SendRequest").Return(&http.Response{
		Status:           "OK",
		StatusCode:       200,
	 }, 1 , nil,
	)

	testService := LoadTestService(mockHTTPClient)


	model :=  models.RequestModel{
		Users:                  1,
		ExecutionTimeInSeconds: 1,
		RampUpTimeInSeconds:    1,
		HTTPRequest:            models.HTTPRequest{
			Method:      "GET",
			URL:         "https://aaa.test.com/health",
		},
	}

	stats, err := testService.RunTest(model)

	mockHTTPClient.AssertExpectations(t)
	mockHTTPClient.AssertCalled(t, "SendRequest")
	assert.Nil(t, err)

	assert.NotNil(t, stats)

}

