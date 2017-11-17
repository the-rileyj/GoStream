package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

type FileInfo struct {
	Name  string
	IsDir bool
	Mode  os.FileMode
}

const (
	filePrefix = "/music/"
	root       = "./music"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "./player.html")
	})
	//Route for public files, aka files in the public folder
	r.GET("/music/:fi", static.Serve("/music", static.LocalFile("music/", true)))
	r.GET("/api/music/:var", func(c *gin.Context) {
		switch c.Param("var") {
		case "list":
			matches, err := filepath.Glob("./music/*.mp3")
			for i, v := range matches {
				matches[i] = strings.TrimPrefix(v, "\\")
				//matches[i] = strings.TrimPrefix(v, "*\\")
				matches[i] = strings.TrimPrefix(v, "music\\")
			}
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(matches)
			c.JSON(200, gin.H{
				"music": matches,
			})
		}
	})
	r.Run(":5000")
}
