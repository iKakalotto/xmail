## XMail

这是一个用来发送邮件的 gRPC 服务，请在 [application.yaml](application.yaml) 中配置邮箱服务器，使用 `env.mail.passwd` 配置环境变量名，用于设置邮箱密码。

安装 protoc-gen-go 插件：
~~~bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
~~~

安装 protoc-gen-go-grpc 插件：
~~~bash
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1
~~~

安装编译器 [protoc](https://github.com/protocolbuffers/protobuf/releases)，编译 proto：
~~~bash
protoc --go_out=. --go-grpc_out=. proto/email.proto
~~~

下载项目依赖的包：
~~~bash
go get
~~~

启动 server:
~~~bash
go run main.go email_grpc_server.go
~~~

测试 client：
~~~bash
go run test.go
~~~
