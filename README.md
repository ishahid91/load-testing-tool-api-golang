# load-testing-tool-api-golang

Load testing API including swagger written in GoLang. API can hosted on any machine and can be used to run Load test on API endpoints.



## Usage

API Json Body
```
{
    "users" : 5,
    "executionTimeInSeconds":20,
    "rampUpTimeInSeconds": 2,
    "httpRequest": {
        "method":"GET",
        "url":"https://nexus.talabat.com/health",
        "headers": {
            "Content-Type": "application/json"
        },
	"queryParams":{
            "param_key":"param_value
        },
	"body": {"json":"value"}
        
    }
}
```
