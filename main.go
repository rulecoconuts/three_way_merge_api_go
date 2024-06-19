package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.POST("/merge", func(c *gin.Context) {

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

	// Get a list of actions needed to perform a merge on an original_file, a_file, and b_file
	router.POST("/merge/actions", func(c *gin.Context) {

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

		c.JSON(http.StatusOK, MergePresentation{Actions: actions, TimeSpentInMicroseconds: elapsed.Microseconds()})

		originalFile.Close()
		v1File.Close()
		v2File.Close()
	})

	router.Run(":8000")
}
