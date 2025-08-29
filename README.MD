# goose
goose 参考 Google API 设计方案，通过定义 proto 文件，快速生成Go语言的Restful服务路由。

# 路由规则
goose 底层基于 [Go 1.22 HTTP Router](https://go.dev/blog/routing-enhancements) 管理http路由。详细定义规则见[Enhanced ServeMux routing](https://github.com/golang/go/issues/61410)

# 优点
1. 基于proto文件，快速生成Restful服务路由
2. 零配置
3. 低依赖
4. 非反射取值，效率非常高
5. 稳定，所有方法经过单元测试
6. 代码简单，良好的代码设计，减少出错

# 安装
```
go install github.com/go-leo/goose/cmd/protoc-gen-goose@latest
```

# Example
```protobuf
syntax = "proto3";
package leo.goose.example.user.v1;
option go_package = "github.com/go-leo/goose/example/user/v1;user";

import "google/api/annotations.proto";

service User {

  // CreateUser 创建用户
  // `POST /v1/user { "name": "Leo" }` | `CreateUserRequest(name: "Leo")`
  rpc CreateUser (CreateUserRequest) returns (CreateUserResponse) {
    option (google.api.http) = {
      post : "/v1/user"
      body : "*"
    };
  }

  // DeleteUser 删除用户
    // `DELETE /v1/user/10000 | `DeleteUserRequest(id: 10000)`
  rpc DeleteUser (DeleteUserRequest) returns (DeleteUserResponse) {
    option (google.api.http) = {
      delete : "/v1/user/{id}"
    };
  }

  // ModifyUser 修改用户
  // `PUT /v1/user/10000 { "name": "Leo" }` | `ModifyUserRequest(id: 10000, name: "Leo")`
  rpc ModifyUser (ModifyUserRequest) returns (ModifyUserResponse) {
    option (google.api.http) = {
      put : "/v1/user/{id}"
      body : "*"
    };
  }

  // UpdateUser 更新用户
  // `PUT /v1/user/10000 { "id": "99999" ,"name": "Leo" }` | `UpdateUserRequest(id: 10000, UserItem(id: 9999, name: "Leo"))`
  rpc UpdateUser (UpdateUserRequest) returns (UpdateUserResponse) {
    option (google.api.http) = {
      patch : "/v1/user/{id}"
      body : "item"
    };
  }

  // GetUser 获取用户
  // `GET /v1/user/10000` | `GetUserRequest(id: 10000)`
  rpc GetUser (GetUserRequest) returns (GetUserResponse) {
    option (google.api.http) = {
      get : "/v1/user/{id}"
    };
  }

  // ListUser 获取用户列表
  // `GET /v1/users?page_num=1&page_size=10` | `ListUserRequest(page_num: 1, page_size: 10)`
  rpc ListUser (ListUserRequest) returns (ListUserResponse) {
    option (google.api.http) = {
      get : "/v1/users"
    };
  }
}

message UserItem {
  int64 id = 1;
  string name = 2;
}

message CreateUserRequest {
  string name = 1;
}

message CreateUserResponse {
  UserItem item = 1;
}

message DeleteUserRequest {
  int64 id = 1;
}

message DeleteUserResponse {
  int64 id = 1;
}

message ModifyUserRequest {
  int64 id = 1;
  string name = 2;
}

message ModifyUserResponse {
    int64 id = 1;
  string name = 2;
}

message UpdateUserRequest {
  int64 id = 1;
  UserItem item = 2;
}

message UpdateUserResponse {
  int64 id = 1;
  UserItem item = 2;
}

message GetUserRequest {
  int64 id = 1;
}

message GetUserResponse {
  UserItem item = 1;
}

message ListUserRequest {
  int64 page_num = 1;
  int64 page_size = 2;
}

message ListUserResponse {
  int64 page_num = 1;
  int64 page_size = 2;
  repeated UserItem list = 3;
}
```
定义了6个接口：分别是
1. 创建用户，POST请求
2. 删除用户，DELETE请求
3. 修改用户，PUT请求
4. 更新用户，PATCH请求
5. 获取用户, GET请求
6. 获取用户列表, GET请求

# 生成路由
执行命令
```
protoc \
--proto_path=. \
--proto_path=../third_party \
--proto_path=../../ \
--go_out=. \
--go_opt=paths=source_relative \
--goose_out=. \
--goose_opt=paths=source_relative \
user/user.proto
```
生成一下文件
```
user
├── user_goose.pb.go
├── user_test.go
├── user.pb.go
└── user.proto
```

# 实现UserService代码：
```go

type MockUserService struct{}

func (m *MockUserService) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	return &CreateUserResponse{Item: &UserItem{Id: 1, Name: req.Name}}, nil
}

func (m *MockUserService) DeleteUser(ctx context.Context, req *DeleteUserRequest) (*DeleteUserResponse, error) {
	return &DeleteUserResponse{Id: req.GetId()}, nil
}

func (m *MockUserService) ModifyUser(ctx context.Context, req *ModifyUserRequest) (*ModifyUserResponse, error) {
	return &ModifyUserResponse{Id: req.GetId(), Name: req.GetName()}, nil
}

func (m *MockUserService) UpdateUser(ctx context.Context, req *UpdateUserRequest) (*UpdateUserResponse, error) {
	return &UpdateUserResponse{Id: req.GetId(), Item: &UserItem{Id: req.GetItem().GetId(), Name: req.GetItem().GetName()}}, nil
}

func (m *MockUserService) GetUser(ctx context.Context, req *GetUserRequest) (*GetUserResponse, error) {
	return &GetUserResponse{Item: &UserItem{Id: req.Id, Name: "test"}}, nil
}

func (m *MockUserService) ListUser(ctx context.Context, req *ListUserRequest) (*ListUserResponse, error) {
	return &ListUserResponse{
		PageNum:  req.GetPageNum(),
		PageSize: req.GetPageSize(),
		List: []*UserItem{
			{Id: 1, Name: "a"},
			{Id: 2, Name: "b"},
		},
	}, nil
}
```

# 创建Server
```go
func main() {
	router := http.NewServeMux()
	router = AppendUserGooseRoute(router, &MockUserService{})
	server := http.Server{Addr: ":8000", Handler: router}
	server.ListenAndServe()
}
```

更多示例见 [example]https://github.com/go-leo/goose/tree/example/user/server

# 子妹项目
* [github.com/go-leo/gonic](https://github.com/go-leo/gonic) 使用 gin-gonic/gin 管理路由。
* [github.com/go-leo/gorilla](https://github.com/go-leo/gorilla) 使用 gorilla/mux 管理路由。
