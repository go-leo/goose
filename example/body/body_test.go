package body

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"google.golang.org/genproto/googleapis/api/httpbody"
	rpchttp "google.golang.org/genproto/googleapis/rpc/http"
	"google.golang.org/protobuf/types/known/emptypb"
)

type MockBodyService struct{}

func (m *MockBodyService) HttpBodyNamedBody(ctx context.Context, request *HttpBodyRequest) (*Response, error) {
	var value BodyRequest
	if err := json.Unmarshal(request.GetBody().GetData(), &value); err != nil {
		return nil, err
	}
	return &Response{Message: value.GetMessage()}, nil
}

func (m *MockBodyService) HttpBodyStarBody(ctx context.Context, request *httpbody.HttpBody) (*Response, error) {
	var value BodyRequest
	if err := json.Unmarshal(request.GetData(), &value); err != nil {
		return nil, err
	}
	return &Response{Message: value.GetMessage()}, nil
}

func (m *MockBodyService) HttpRequest(ctx context.Context, request *rpchttp.HttpRequest) (*Response, error) {
	var value BodyRequest
	if err := json.Unmarshal(request.GetBody(), &value); err != nil {
		return nil, err
	}
	return &Response{Message: value.GetMessage()}, nil
}

func (m *MockBodyService) NamedBody(ctx context.Context, request *NamedBodyRequest) (*Response, error) {
	return &Response{Message: string(request.GetBody().GetMessage())}, nil
}

func (m *MockBodyService) NonBody(ctx context.Context, request *emptypb.Empty) (*Response, error) {
	return &Response{Message: "NonBody"}, nil
}

func (m *MockBodyService) StarBody(ctx context.Context, request *BodyRequest) (*Response, error) {
	return &Response{Message: request.GetMessage()}, nil
}

func runServer(server *http.Server) {
	router := http.NewServeMux()
	router = AppendBodyGooseRoute(router, &MockBodyService{})
	server.Addr = ":28080"
	server.Handler = router
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func TestStarBody(t *testing.T) {
	router := http.NewServeMux()
	router = AppendBodyGooseRoute(router, &MockBodyService{})
	server := httptest.NewServer(router)
	url := server.URL + "/v1/star/body"
	contentType := "application/json"
	payload := strings.NewReader(`{"message": "hello"}`)
	res, err := http.Post(url, contentType, payload)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != `{"message":"hello"}` {
		t.Fatal("body is not equal")
	}
}

func TestNamedBody(t *testing.T) {
	router := http.NewServeMux()
	router = AppendBodyGooseRoute(router, &MockBodyService{})
	server := httptest.NewServer(router)
	url := server.URL + "/v1/star/body"
	contentType := "application/json"
	payload := strings.NewReader(`{"message": "hello"}`)
	res, err := http.Post(url, contentType, payload)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != `{"message":"hello"}` {
		t.Fatal("body is not equal")
	}
}

func TestNonBody(t *testing.T) {
	router := http.NewServeMux()
	router = AppendBodyGooseRoute(router, &MockBodyService{})
	server := httptest.NewServer(router)
	url := server.URL + "/v1/star/body"
	contentType := "application/json"
	payload := strings.NewReader(`{"message": "hello"}`)
	res, err := http.Post(url, contentType, payload)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != `{"message":"hello"}` {
		t.Fatal("body is not equal")
	}
}

func TestHttpBodyStarBody(t *testing.T) {
	router := http.NewServeMux()
	router = AppendBodyGooseRoute(router, &MockBodyService{})
	server := httptest.NewServer(router)
	url := server.URL + "/v1/star/body"
	contentType := "application/json"
	payload := strings.NewReader(`{"message": "hello"}`)
	res, err := http.Post(url, contentType, payload)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != `{"message":"hello"}` {
		t.Fatal("body is not equal")
	}
}

func TestHttpBodyNamedBody(t *testing.T) {
	router := http.NewServeMux()
	router = AppendBodyGooseRoute(router, &MockBodyService{})
	server := httptest.NewServer(router)
	url := server.URL + "/v1/star/body"
	contentType := "application/json"
	payload := strings.NewReader(`{"message": "hello"}`)
	res, err := http.Post(url, contentType, payload)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != `{"message":"hello"}` {
		t.Fatal("body is not equal")
	}
}

func TestHttpRequest(t *testing.T) {
	router := http.NewServeMux()
	router = AppendBodyGooseRoute(router, &MockBodyService{})
	server := httptest.NewServer(router)
	url := server.URL + "/v1/star/body"
	contentType := "application/json"
	payload := strings.NewReader(`{"message": "hello"}`)
	res, err := http.Post(url, contentType, payload)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != `{"message":"hello"}` {
		t.Fatal("body is not equal")
	}
}
