package user

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	gin "github.com/gin-gonic/gin"
)

// ---- Mock Service ----

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

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router = AppendUserGonicRoute(router, &MockUserService{})
	server := http.Server{Addr: ":8000", Handler: router}
	server.ListenAndServe()
}

// ---- Test Cases ----

func setupServer() *httptest.Server {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router = AppendUserGonicRoute(router, &MockUserService{})
	return httptest.NewServer(router)
}

func TestCreateUser(t *testing.T) {
	server := setupServer()
	defer server.Close()

	payload := []byte(`{"name":"alice"}`)
	resp, err := http.Post(server.URL+"/v1/user", "application/json", bytes.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	expected := `{"item":{"id":"1","name":"alice"}}`
	if strings.ReplaceAll(string(body), " ", "") != strings.ReplaceAll(expected, " ", "") {
		t.Fatalf("body is not equal: got %s, want %s", string(body), expected)
	}
}

func TestDeleteUser(t *testing.T) {
	server := setupServer()
	defer server.Close()

	req, _ := http.NewRequest(http.MethodDelete, server.URL+"/v1/user/123", nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status: %d", resp.StatusCode)
	}
	body, _ := io.ReadAll(resp.Body)
	expected := `{"id":"123"}`
	if strings.ReplaceAll(string(body), " ", "") != strings.ReplaceAll(expected, " ", "") {
		t.Fatalf("body is not equal: got %s, want %s", string(body), expected)
	}
}

func TestModifyUser(t *testing.T) {
	server := setupServer()
	defer server.Close()

	payload := []byte(`{"id":123,"name":"bob"}`)
	req, _ := http.NewRequest(http.MethodPut, server.URL+"/v1/user/123", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status: %d", resp.StatusCode)
	}
	body, _ := io.ReadAll(resp.Body)
	expected := `{"id":"123", "name":"bob"}`
	if strings.ReplaceAll(string(body), " ", "") != strings.ReplaceAll(expected, " ", "") {
		t.Fatalf("body is not equal: got %s, want %s", string(body), expected)
	}
}

func TestUpdateUser(t *testing.T) {
	server := setupServer()
	defer server.Close()
	payload := []byte(`{"id":567,"name":"bob"}`)
	req, _ := http.NewRequest(http.MethodPatch, server.URL+"/v1/user/123", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status: %d", resp.StatusCode)
	}
	body, _ := io.ReadAll(resp.Body)
	expected := `{"id":"123", "item":{"id":"567", "name":"bob"}}`
	if strings.ReplaceAll(string(body), " ", "") != strings.ReplaceAll(expected, " ", "") {
		t.Fatalf("body is not equal: got %s, want %s", string(body), expected)
	}
}

func TestGetUser(t *testing.T) {
	server := setupServer()
	defer server.Close()

	resp, err := http.Get(server.URL + "/v1/user/123")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	expected := `{"item":{"id":"123","name":"test"}}`
	if strings.ReplaceAll(string(body), " ", "") != strings.ReplaceAll(expected, " ", "") {
		t.Fatalf("body is not equal: got %s, want %s", string(body), expected)
	}
}

func TestListUser(t *testing.T) {
	server := setupServer()
	defer server.Close()

	resp, err := http.Get(server.URL + "/v1/users?page_num=1&page_size=10")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	expected := `{"pageNum":"1", "pageSize":"10", "list":[{"id":"1", "name":"a"}, {"id":"2", "name":"b"}]}`
	if strings.ReplaceAll(string(body), " ", "") != strings.ReplaceAll(expected, " ", "") {
		t.Fatalf("body is not equal: got %s, want %s", string(body), expected)
	}
}
