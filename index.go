package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

type FileInfo struct {
	Length string `json:"length"`
	Name   string `json:"name"`
}

const (
	filePrefix = "/music/"
	root       = "./music"
)

func main() {
	d := []FileInfo{}
	if r, err := os.Open("data.json"); err == nil {
		if b, err := ioutil.ReadAll(bufio.NewReader(r)); err == nil {
			json.Unmarshal(b, &d)
		} else {
			fmt.Printf("ERROR 1:\n%v\n", err)
		}
	} else {
		fmt.Printf("ERROR 2:\n%v\n", err)
	}
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "./player.html")
	})
	//Route for public files, aka files in the public folder
	r.GET("/music/:fi", static.Serve("/music", static.LocalFile("music/", true)))
	r.GET("/api/music/:var", func(c *gin.Context) {
		switch c.Param("var") {
		case "list":
			/*matches, err := filepath.Glob("./music/*.mp3")
			for i, v := range matches {
				matches[i] = strings.TrimPrefix(v, "\\")
				//matches[i] = strings.TrimPrefix(v, "*\\")
				matches[i] = strings.TrimPrefix(v, "music\\")
			}
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(matches)*/
			fmt.Println(d)
			c.JSON(200, gin.H{
				"list": d,
			})
		}
	})
	r.Run(":5000")
}
