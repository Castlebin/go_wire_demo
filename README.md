# Golang依赖注入框架wire使用详解
https://zhuanlan.zhihu.com/p/338926709

## 1. 初始化项目
```shell
mkdir go_wire_demo
cd go_wire_demo
go mod init go_wire_demo
```

### 1.1. 安装 wire 命令行工具
```shell
# 安装命令 （注意，将 GOPATH/bin 加入 PATH）
go install github.com/google/wire/cmd/wire
```

### 1.2. 项目安装依赖
```shell
go get -u github.com/google/wire
```

## 2. Quick Start
### 2.1. 新建 main.go 文件
```go
package main

import "fmt"

type Message struct {
	msg string
}
type Greeter struct {
	Message Message
}
type Event struct {
	Greeter Greeter
}

// NewMessage Message的构造函数
func NewMessage(msg string) Message {
	return Message{
		msg: msg,
	}
}

// NewGreeter Greeter构造函数
func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

// NewEvent Event构造函数
func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

func (g Greeter) Greet() Message {
	return g.Message
}

/**
项目需要 3 个对象，对象关系如下：
Event (
	Greeter : Greeter (
		Message : Message (
			msg : string
		)
	)
)
*/

// 使用wire前
func main() {
	message := NewMessage("hello world")
	greeter := NewGreeter(message)
	event := NewEvent(greeter)

	event.Start()
}
```
