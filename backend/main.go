package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
)

var ginLambda *ginadapter.GinLambda

func generateEngine() *gin.Engine {
	engine := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowWildcard = true
	// corsConfig.AllowOrigins = []string{"https://ofejiro.com", "https://*.ofejiro.com", "*"}
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowHeaders("Origin")

	engine.Use(cors.New(corsConfig))

	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"*"},
	// 	AllowHeaders:     []string{"Origin"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// }))

	engine.POST("/merge", func(c *gin.Context) {

		form, _ := c.MultipartForm()
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		fmt.Println(form.File)

		originalFileHeader := form.File["original_file"][0]
		v1FileHeader := form.File["v1_file"][0]
		v2FileHeader := form.File["v2_file"][0]

		originalFile, _ := originalFileHeader.Open()
		v1File, _ := v1FileHeader.Open()
		v2File, _ := v2FileHeader.Open()

		c.Header("Content-Type", "text/plain")
		c.Header("Content-Disposition", "attachment; filename=merged.txt")

		err := MergeFromMultipartFileIntoWriter(originalFile, v1File, v2File, c.Writer)

		if err != nil {
			c.AbortWithError(500, err)
		}

		originalFile.Close()
		v1File.Close()
		v2File.Close()
	})

	// Get a list of actions needed to perform a merge on an original_file, v1_file, and v2_file
	engine.POST("/merge/actions", func(c *gin.Context) {

		form, _ := c.MultipartForm()
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		originalFileHeader := form.File["original_file"][0]
		v1FileHeader := form.File["v1_file"][0]
		v2FileHeader := form.File["v2_file"][0]

		originalFile, _ := originalFileHeader.Open()
		v1File, _ := v1FileHeader.Open()
		v2File, _ := v2FileHeader.Open()

		c.Header("Content-Type", "application/json")
		// c.Header("Content-Disposition", "attachment; filename=merged.txt")

		start := time.Now()

		// Get a list of actions needed to merge 3 versions of a file
		actions, err := GetMergeActionsFromMultipartFiles(originalFile, v1File, v2File)

		elapsed := time.Since(start)
		fmt.Println(actions)

		if err != nil {
			c.AbortWithError(500, err)
		}

		// We are measuring time in microseconds because the average time measured during tests ended up as microseconds.
		// even if we obtain a time in a higher unit, the requester can always perform a conversion
		c.JSON(http.StatusOK, MergePresentation{Actions: actions, TimeSpentInMicroseconds: elapsed.Microseconds()})

		originalFile.Close()
		v1File.Close()
		v2File.Close()
	})

	return engine
}

func init() {
	ginLambda = ginadapter.New(generateEngine())
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, request)
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	lambda.Start(Handler)
}
