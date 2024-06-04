package main

import "github.com/gin-gonic/gin"

func main() {
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.POST("/merge", func(c *gin.Context) {

		form, _ := c.MultipartForm()

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

	router.Run(":8000")
}
