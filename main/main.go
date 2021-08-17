package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "load-test-tool/main/docs"
	"load-test-tool/models"
	"load-test-tool/services"
	"net/http"
)


func health(c *gin.Context) {
	c.IndentedJSON(http.StatusOK,"Health Check Ok")
}


// RunLoadTest godoc
// @Summary Run Load Test on specified Url
// @Description Run Load Test on specified Url
// @Accept  json
// @Produce  json
// @Param request body string true "Request Model"
// @Success 200 {object} string
// @Router /runloadtest [post]
func runLoadTest(c *gin.Context) {
	var request models.RequestModel
	if err := c.BindJSON(&request); err != nil {
		return
	}
	// Get a greeting message and print it.

	var loadService services.LoadTestModel
	result := loadService.RunTest(request)
	c.IndentedJSON(http.StatusOK, result)
}



// @title Load Test API
// @version 1.0
// @description Load Test tool
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:80
// @BasePath /
func main() {
	router := gin.Default()
	router.GET("/health", health)
	router.POST("/runloadtest", runLoadTest)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run("localhost:80")
}