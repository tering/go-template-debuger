# Golang template Debuger

在 gopher 的日常开发工作中，时长会遇到写模板的需求。但大多数的 gopher（包括我在内）都不太了解 go template 的语法规则。
为了降低认知负担和试错成本，于是我开发了这个调试 Go 模板的工具。

## 使用说明

### 源码方式使用
1. 下载代码并进入本项目目录
2. 在本项目文件夹下编写模板文件： `a.html`
3. 在本项目文件夹下编写模板数据： `a.json`
4. 启动服务，本工具假设你已安装了 Go 开发环境
```bash
go run main.go
```
1. 访问 http://localhost:8080/view/a.html 即可看到渲染后的结果

### 安装命令行方式使用
1. 安装本工具
```bash
# 从 gitee 安装
go install gitee.com/tering/go-template-debuger

# 或从 github 安装
go install github.com/tering/go-template-debuger
```
2. 在模板文件所在目录下执行命令： `go-template-debuger`
3. 在模板文件所在目录下创建同名的 json 文件。例如： 模板文件为`a.email`, 则数据文件为`a.json`
4. 访问 http://localhost:8080/view/<模板文件名> 即可看到渲染后的结果

## 模板语法

具体语法规则请参考 [Go 模板语法](https://golang.org/pkg/text/template/#pkg-overview)。
