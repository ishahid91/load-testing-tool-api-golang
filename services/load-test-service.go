package services

import (
	"load-test-tool/models"
	"load-test-tool/utilities"
	"log"
	"time"
)


type LoadTestModel struct {
	Stats models.LoadTestStats
}

type ILoadTestService interface {
	RunTest(model models.RequestModel) (models.LoadTestStats, error)
	CalculateStats(stats models.CurrentStats) error
}

func LoadTestService(httpClient utilities.IHTTPClientService) ILoadTestService {
	_httpClient = httpClient
	return &LoadTestModel{}
}

var _httpClient utilities.IHTTPClientService

func (i *LoadTestModel) RunTest(request models.RequestModel) (models.LoadTestStats, error) {

	log.Printf("Load Test started for, %v \n!", request.HTTPRequest.URL)

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
			return i.Stats, nil
		case userActive := <- threadsCh:
			log.Printf("User %d is started", userActive)
			go activeUserCalls(userActive,request, activeUsersCh, loadResultCh)
		case httpResult := <- loadResultCh:
			//log.Printf("%v",httpResult)
			i.CalculateStats(httpResult)
		}
	}
}

func activeUserCalls(userCount int64,request models.RequestModel , ch chan bool , loadResultCh chan models.CurrentStats) error {

	for {
		select {
		case <- ch:
			log.Printf("User %d is stopping", userCount)
			return nil
		default:
			//log.Printf("User %d is running", userCount)
			sendRequests(request, loadResultCh)
		}
	}
}


func createThreads(request models.RequestModel, ch chan <- int64) error {
	waitTime := request.RampUpTimeInSeconds / request.Users
	var i int64
	for i = 0; i < request.Users; i++ {
		ch <- i
		time.Sleep(time.Duration(waitTime) * time.Second)
	}
	return nil
}

func sendRequests(request models.RequestModel,  loadResultCh chan <- models.CurrentStats) error {
	var metric models.CurrentStats

	response, responseTimeInMillieconds, err := _httpClient.SendRequest(request.HTTPRequest.Method, request.HTTPRequest.URL, request.HTTPRequest.QueryParams, request.HTTPRequest.Headers,
		request.HTTPRequest.Body)
	if err != nil {
		log.Print("Error while making HTTP request :", err.Error())
		return err
	}

	metric.ResponseTime = responseTimeInMillieconds
	statusOK := response.StatusCode >= 200 && response.StatusCode < 300
	//log.Printf("Response %v, %v", response.Status, responseTimeInMillieconds)
	if statusOK {
		metric.Error = false
		loadResultCh <- metric
		return nil
	}

	metric.Error = true
	loadResultCh <- metric
    return err
}

func (i *LoadTestModel)CalculateStats(stat models.CurrentStats) error{
	i.Stats.TotalRequests++
	if stat.Error {
		i.Stats.ErrorCount++
	}
	i.Stats.AverageResponseTime = ((i.Stats.AverageResponseTime * (i.Stats.TotalRequests -1)) + stat.ResponseTime) / i.Stats.TotalRequests
	if stat.ResponseTime > i.Stats.MaxResponseTime {
		i.Stats.MaxResponseTime = stat.ResponseTime
	}
	return nil
}