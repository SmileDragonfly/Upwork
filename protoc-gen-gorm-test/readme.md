# Generate protobuf files
1. Install protoc: https://github.com/protocolbuffers/protobuf/releases/tag/v25.3
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
```
2. Install protoc-gen-go-grpc
```bash
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```
3. Install protoc-gen-gorm
```bash
go install github.com/infobloxopen/protoc-gen-gorm@latest
```
4. Generate model files
``` bash
protoc --proto_path="C:\Users\datdo\go\pkg\mod\github.com\infobloxopen\protoc-gen-gorm@v1.1.2\proto\options;." --go_out=.\user.pb\ --go_opt=paths=source_relative --go-grpc_out=.\user.pb\ --go-grpc_opt=paths=source_relative user.proto --gorm_out=.
```
*Notice: Change --proto_path to your local path of protoc-gen-gorm*
5. Notice to generate CRUD correctly
- For CRUD methods to be generated correctly you need to follow specific conventions:
  - Request messages for Create and Update methods should have an Ormable Type in a field named payload, for Read and Delete methods an id field is required. Nothing is required in the List request.
  - Response messages for Create, Read, and Update require an Ormable Type in a field named result and for List a repeated Ormable Type named results.
  - Delete methods require the (gorm.method).object_type option to indicate which Ormable Type it should delete, and has no response type requirements.
6. For this project
- Run main: go run .
- Run test CRUD: cd user.pb
  - Create: go test -run TestUserServiceDefaultServer_Create
  - Read: go test -run TestUserServiceDefaultServer_Read
  - Update: go test -run TestUserServiceDefaultServer_Update
  - Delete: go test -run TestUserServiceDefaultServer_Delete
   