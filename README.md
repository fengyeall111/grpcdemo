# grpc 入门demo

- 编写proto文件，pb/guide.proto
- 编译proto文件, 
protoc --go_out=. .\pb\guide.proto
protoc --go-grpc_out=. .\pb\guide.proto
- 编写server,
1. s = grpc.newServer()
2. grpc.registerService(&server{})
3. 实现proto中定义的服务
4. s.serve(tcpconn)
- 编写client
1. tcpconn
2. newprpcClient(conn)
3. 按照grpc提供的接口调用服务
- 运行客户端和服务端
go run grpcdemo/server/guide.go
- go run grpcdemo/client/guide.go
