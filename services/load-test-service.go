package services

import (
	"bytes"
	"encoding/json"
	"load-test-tool/models"
	"load-test-tool/utilities"
	"log"
	"time"
)


type LoadTestModel struct {
	Stats models.LoadTestStats
}

type ILoadTestService interface {
	RunTest(model models.RequestModel) models.LoadTestStats
	CalculateStats(stats models.CurrentStats)

}

func LoadTestService() ILoadTestService {
	return &LoadTestModel{}
}

var httpClient utilities.IHTTPClientService = utilities.HTTPClientService()

func (i *LoadTestModel) RunTest(request models.RequestModel) models.LoadTestStats {

	log.Printf("Load Test started for, %v \n!", request.HTTPRequest.URL)


	//i.HTTPClient.Initialise()

	threadsCh := make(chan int64)
	activeUsersCh := make(chan bool)
	go createThreads(request, threadsCh)
	timeoutCh := time.After(time.Duration(request.ExecutionTimeInSeconds) * time.Second)

	loadResultCh := make(chan models.CurrentStats)

	for {
		select {
		case <- timeoutCh:
			log.Printf("Execution Completed !!")
			close(activeUsersCh)
			return i.Stats
		case userActive := <- threadsCh:
			log.Printf("User %d is started", userActive)
			go activeUserCalls(userActive,request, activeUsersCh, loadResultCh)
		case httpResult := <- loadResultCh:
			//log.Printf("%v",httpResult)
			i.CalculateStats(httpResult)
		}
	}
}

func activeUserCalls(userCount int64,request models.RequestModel , ch chan bool , loadResultCh chan models.CurrentStats) {

	for {
		select {
		case <- ch:
			log.Printf("User %d is stopping", userCount)
			return
		default:
			//log.Printf("User %d is running", userCount)
			sendRequests(request, loadResultCh)
		}
	}
}


func createThreads(request models.RequestModel, ch chan <- int64) {
	waitTime := request.RampUpTimeInSeconds / request.Users
	var i int64
	for i = 0; i < request.Users; i++ {
		ch <- i
		time.Sleep(time.Duration(waitTime) * time.Second)
	}
}

func sendRequests(request models.RequestModel,  loadResultCh chan <- models.CurrentStats) {
	var metric models.CurrentStats
	payloadBytes, err := json.Marshal(request.HTTPRequest.Body)
	if err != nil {
		log.Print("Unable to convert payload to bytes : ", err.Error())
	}
	payload := bytes.NewReader(payloadBytes)
	response, responseTimeInMillieconds, err := httpClient.SendRequest(request.HTTPRequest.Method, request.HTTPRequest.URL, nil, request.HTTPRequest.Headers,
		payload)
	if err != nil {
		log.Print("Error while making HTTP request :", err.Error())
	}

	metric.ResponseTime = responseTimeInMillieconds
	statusOK := response.StatusCode >= 200 && response.StatusCode < 300
	//log.Printf("Response %v, %v", response.Status, responseTimeInMillieconds)
	if statusOK {
		metric.Error = false
		loadResultCh <- metric
		return
	}

	metric.Error = true
	loadResultCh <- metric
    return
}

func (i *LoadTestModel)CalculateStats(stat models.CurrentStats) {
	i.Stats.TotalRequests++
	if stat.Error {
		i.Stats.ErrorCount++
	}
	i.Stats.AverageResponseTime = ((i.Stats.AverageResponseTime * (i.Stats.TotalRequests -1)) + stat.ResponseTime) / i.Stats.TotalRequests
	if stat.ResponseTime > i.Stats.MaxResponseTime {
		i.Stats.MaxResponseTime = stat.ResponseTime
	}
}