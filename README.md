# go-arduino-led64-server

### 开发平台

- Go 1.21

### 功能结构

在本项目中，后端向前端提供以下三个接口：

- /config     GET        返回当前连接LED设备的板载模式
- /status     GET        返回当前连接LED设备的当前状态
- /set           POST     更改当前连接LED设备的模式与状态

为了实现以上接口，后端程序包含以下几个模块：

- main.go	主函数入口，管理接口入口与跨域处理
- http.go          实现接口处理函数
- json.go          解析前端传递的json信息
- udp.go           封装函数，实现与受控LED设备的udp通信
