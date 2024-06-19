package main

import (
	"fmt"

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
		aFileHeader := form.File["a_file"][0]
		bFileHeader := form.File["b_file"][0]

		originalFile, _ := originalFileHeader.Open()
		aFile, _ := aFileHeader.Open()
		bFile, _ := bFileHeader.Open()

		c.Header("Content-Type", "text/plain")
		c.Header("Content-Disposition", "attachment; filename=merged.txt")

		err := MergeFromMultipartFileIntoWriter(originalFile, aFile, bFile, c.Writer)

		if err != nil {
			c.AbortWithError(500, err)
		}

		originalFile.Close()
		aFile.Close()
		bFile.Close()
	})

	router.POST("/merge/actions", func(c *gin.Context) {

		form, _ := c.MultipartForm()
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		originalFileHeader := form.File["original_file"][0]
		aFileHeader := form.File["a_file"][0]
		bFileHeader := form.File["b_file"][0]

		originalFile, _ := originalFileHeader.Open()
		aFile, _ := aFileHeader.Open()
		bFile, _ := bFileHeader.Open()

		c.Header("Content-Type", "application/json")
		// c.Header("Content-Disposition", "attachment; filename=merged.txt")

		actions, err := GetMergeActionsFromMultipartFiles(originalFile, aFile, bFile)
		fmt.Println(actions)

		if err != nil {
			c.AbortWithError(500, err)
		}

		c.JSON(200, actions)

		originalFile.Close()
		aFile.Close()
		bFile.Close()
	})

	router.Run(":8000")
}
