package response_body

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	rpchttp "google.golang.org/genproto/googleapis/rpc/http"
)

// ---- Mock Service ----

type MockResponseBodyService struct{}

func (m *MockResponseBodyService) OmittedResponse(ctx context.Context, req *Request) (*Response, error) {
	return &Response{Message: "omitted"}, nil
}

func (m *MockResponseBodyService) StarResponse(ctx context.Context, req *Request) (*Response, error) {
	return &Response{Message: "star"}, nil
}

func (m *MockResponseBodyService) NamedResponse(ctx context.Context, req *Request) (*NamedBodyResponse, error) {
	return &NamedBodyResponse{
		Body: &NamedBodyResponse_Body{Message: "named"},
	}, nil
}

func (m *MockResponseBodyService) HttpBodyResponse(ctx context.Context, req *Request) (*httpbody.HttpBody, error) {
	return &httpbody.HttpBody{
		ContentType: "text/plain",
		Data:        []byte("httpbody"),
	}, nil
}

func (m *MockResponseBodyService) HttpBodyNamedResponse(ctx context.Context, req *Request) (*NamedHttpBodyResponse, error) {
	return &NamedHttpBodyResponse{
		Body: &httpbody.HttpBody{
			ContentType: "application/json",
			Data:        []byte(`{"message":"httpbodynamed"}`),
		},
	}, nil
}

func (m *MockResponseBodyService) HttpResponse(ctx context.Context, req *Request) (*rpchttp.HttpResponse, error) {
	return &rpchttp.HttpResponse{
		Status: 200,
		Body:   []byte("httpresponse"),
	}, nil
}

// ---- Test Cases ----

func TestOmittedResponse(t *testing.T) {
	router := http.NewServeMux()
	router = AppendResponseBodyGooseRoute(router, &MockResponseBodyService{})
	server := httptest.NewServer(router)
	defer server.Close()

	url := server.URL + "/v1/omitted/response"
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	expected := `{"message":"omitted"}`
	if string(body) != expected {
		t.Fatalf("body is not equal: got %s, want %s", string(body), expected)
	}
}

func TestStarResponse(t *testing.T) {
	router := http.NewServeMux()
	router = AppendResponseBodyGooseRoute(router, &MockResponseBodyService{})
	server := httptest.NewServer(router)
	defer server.Close()

	url := server.URL + "/v1/star/response"
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	expected := `{"message":"star"}`
	if string(body) != expected {
		t.Fatalf("body is not equal: got %s, want %s", string(body), expected)
	}
}

func TestNamedResponse(t *testing.T) {
	router := http.NewServeMux()
	router = AppendResponseBodyGooseRoute(router, &MockResponseBodyService{})
	server := httptest.NewServer(router)
	defer server.Close()

	url := server.URL + "/v1/named/response"
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	expected := `{"message":"named"}`
	if string(body) != expected {
		t.Fatalf("body is not equal: got %s, want %s", string(body), expected)
	}
}

func TestHttpBodyResponse(t *testing.T) {
	router := http.NewServeMux()
	router = AppendResponseBodyGooseRoute(router, &MockResponseBodyService{})
	server := httptest.NewServer(router)
	defer server.Close()

	url := server.URL + "/v1/http/body/omitted/response"
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	expected := "httpbody"
	if string(body) != expected {
		t.Fatalf("body is not equal: got %s, want %s", string(body), expected)
	}
}

func TestHttpBodyNamedResponse(t *testing.T) {
	router := http.NewServeMux()
	router = AppendResponseBodyGooseRoute(router, &MockResponseBodyService{})
	server := httptest.NewServer(router)
	defer server.Close()

	url := server.URL + "/v1/http/body/named/response"
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	expected := `{"message":"httpbodynamed"}`
	if string(body) != expected {
		t.Fatalf("body is not equal: got %s, want %s", string(body), expected)
	}
}

func TestHttpResponse(t *testing.T) {
	router := http.NewServeMux()
	router = AppendResponseBodyGooseRoute(router, &MockResponseBodyService{})
	server := httptest.NewServer(router)
	defer server.Close()

	url := server.URL + "/v1/http/response"
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	expected := "httpresponse"
	if string(body) != expected {
		t.Fatalf("body is not equal: got %s, want %s", string(body), expected)
	}
}
