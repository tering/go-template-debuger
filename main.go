package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

func main() {
	r := gin.Default()

	// 访问根目录时，重定向到Gitee仓库地址
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusSeeOther, "https://gitee.com/tering/go-template-debuger")
	})

	r.GET("/:fileName", func(c *gin.Context) {
		fileName := c.Param("fileName")

		// 读取模板文件
		t, err := template.ParseFiles(fmt.Sprintf("./template/%s.html", fileName))
		if err != nil {
			c.String(500, err.Error())
			return
		}

		// 读取需要传入模板的JSON数据
		jsonData, err := os.ReadFile(fmt.Sprintf("./template/%s.json", fileName))
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

		// 将数据写入文件
		err = t.Execute(c.Writer, data)
		if err != nil {
			c.String(500, err.Error())
		}
	})
	r.Run()
}
