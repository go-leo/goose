# goose
goose 通过定义 proto 文件，快速生成基于 Go 1.22 HTTP Router 的Restful服务路由和客户端代码。

# 优点
1. 基于proto文件，快速生成Restful服务路由和客户端
2. 零配置
3. 低依赖
4. 非反射取值，效率非常高
5. 稳定，所有方法经过单元测试
6. 代码简单，良好的代码设计，减少出错

# 安装
```
go install github.com/go-leo/goose/cmd/protoc-gen-goose@latest
```

# 路由规则
goose 底层基于 [Go 1.22 HTTP Router](https://go.dev/blog/routing-enhancements) 管理http路由。
## Google Http API Annotations
[Google Http API Annotations](https://github.com/googleapis/googleapis/blob/master/google/api/annotations.proto)是google定义的可以在proto文件中使用的http路由规则注解。

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
## 1. CreateUser - 创建用户
- **HTTP 方法**: POST
- **路由路径**: `/v1/user`
- **请求体**: 整个请求消息 (`body: "*"`)
- **示例**:
  ```
  POST /v1/user
  {
    "name": "Leo"
  }
  ```

## 2. DeleteUser - 删除用户
- **HTTP 方法**: DELETE
- **路由路径**: `/v1/user/{id}`
- **路径参数**: `id` 字段从请求消息中提取
- **示例**:
  ```
  DELETE /v1/user/10000
  ```

## 3. ModifyUser - 修改用户
- **HTTP 方法**: PUT
- **路由路径**: `/v1/user/{id}`
- **路径参数**: `id` 字段从请求消息中提取
- **请求体**: 整个请求消息 (`body: "*"`)
- **示例**:
  ```
  PUT /v1/user/10000
  {
    "name": "Leo"
  }
  ```

## 4. UpdateUser - 更新用户
- **HTTP 方法**: PATCH
- **路由路径**: `/v1/user/{id}`
- **路径参数**: `id` 字段从请求消息中提取
- **请求体**: 仅 `item` 字段 (`body: "item"`)
- **示例**:
  ```
  PATCH /v1/user/10000
  {
    "item": {
      "id": 9999,
      "name": "Leo"
    }
  }
  ```

## 5. GetUser - 获取用户
- **HTTP 方法**: GET
- **路由路径**: `/v1/user/{id}`
- **路径参数**: `id` 字段从请求消息中提取
- **示例**:
  ```
  GET /v1/user/10000
  ```

## 6. ListUser - 获取用户列表
- **HTTP 方法**: GET
- **路由路径**: `/v1/users`
- **查询参数**: `page_num` 和 `page_size` 字段会自动转换为查询参数
- **示例**:
  ```
  GET /v1/users?page_num=1&page_size=10
  ```

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
