package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	file, _ := c.FormFile("file")
	fmt.Printf("file name: %s\n", file.Filename)

	err := c.SaveUploadedFile(file, "aFile")
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	c.String(http.StatusOK, "ok")
}
