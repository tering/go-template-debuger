package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

func main() {
	r := gin.Default()
	r.StaticFS("/static", http.Dir("./"))

	// 访问根目录时，重定向到Gitee仓库地址
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/static")
	})

	r.GET("view/:fileName", func(c *gin.Context) {
		fileName := c.Param("fileName")

		// 加载模板文件
		r.LoadHTMLFiles(fileName)

		// 读取需要传入模板的JSON数据
		jsonData, err := os.ReadFile(fileName[:strings.LastIndex(fileName, ".")] + ".json")
		if err != nil {
			c.String(500, err.Error())
			return
		}

		// 解析需要传入模板的JSON数据
		var data interface{}
		err = jsoniter.Unmarshal(jsonData, &data)
		if err != nil {
			c.String(500, err.Error())
			return
		}

		// 响应模板文件
		c.HTML(200, fileName, data)
	})
	r.Run()
}
