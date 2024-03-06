# Generate protobuf files
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
```
```bash
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```
``` bash
protoc --go_out=.\user\ --go_opt=paths=source_relative --go-grpc_out=.\user\ --go-grpc_opt=paths=source_relative user.proto
``` 

