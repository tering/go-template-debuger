# Golang template Debuger

在 gopher 的日常开发工作中，时长会遇到写模板的需求。但大多数的 gopher（包括我在内）都不太了解 go template 的语法规则。
为了降低认知负担和试错成本，于是我开发了这个调试 Go 模板的工具。

## 使用说明

1. 下载代码
```bash
git clone https://gitee.com/tering/go-template-debuger.git
cd go-template-debuger
```
2. 在 template 文件夹下编写模板文件： `a.html`
3. 在 template 文件夹下编写模板数据： `a.json`
4. 启动服务，本工具假设你已安装了 Go 开发环境
```bash
go run main.go
```
5. 访问 http://localhost:8080/a 即可看到渲染后的结果

## 模板语法

具体语法规则请参考 [Go 模板语法](https://golang.org/pkg/text/template/#pkg-overview)。
