package main

import (
	"context"
	"net/http"

	"github.com/go-leo/goose/example/user"
)

type MockUserService struct{}

func (m *MockUserService) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	return &user.CreateUserResponse{Item: &user.UserItem{Id: 1, Name: req.Name}}, nil
}

func (m *MockUserService) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (*user.DeleteUserResponse, error) {
	return &user.DeleteUserResponse{Id: req.GetId()}, nil
}

func (m *MockUserService) ModifyUser(ctx context.Context, req *user.ModifyUserRequest) (*user.ModifyUserResponse, error) {
	return &user.ModifyUserResponse{Id: req.GetId(), Name: req.GetName()}, nil
}

func (m *MockUserService) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (*user.UpdateUserResponse, error) {
	return &user.UpdateUserResponse{Id: req.GetId(), Item: &user.UserItem{Id: req.GetItem().GetId(), Name: req.GetItem().GetName()}}, nil
}

func (m *MockUserService) GetUser(ctx context.Context, req *user.GetUserRequest) (*user.GetUserResponse, error) {
	return &user.GetUserResponse{Item: &user.UserItem{Id: req.Id, Name: "test"}}, nil
}

func (m *MockUserService) ListUser(ctx context.Context, req *user.ListUserRequest) (*user.ListUserResponse, error) {
	return &user.ListUserResponse{
		PageNum:  req.GetPageNum(),
		PageSize: req.GetPageSize(),
		List: []*user.UserItem{
			{Id: 1, Name: "a"},
			{Id: 2, Name: "b"},
		},
	}, nil
}

func main() {
	router := http.NewServeMux()
	router = user.AppendUserGooseRoute(router, &MockUserService{})
	server := http.Server{Addr: ":8000", Handler: router}
	server.ListenAndServe()
}
