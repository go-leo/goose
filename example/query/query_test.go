package query

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	gin "github.com/gin-gonic/gin"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// ---- Mock Services ----

type MockBoolQueryService struct{}

func (m *MockBoolQueryService) BoolQuery(ctx context.Context, req *BoolQueryRequest) (*httpbody.HttpBody, error) {
	data, err := protojson.Marshal(req)
	if err != nil {
		return nil, err
	}
	return &httpbody.HttpBody{Data: data}, nil
}

type MockInt32QueryService struct{}

func (m *MockInt32QueryService) Int32Query(ctx context.Context, req *Int32QueryRequest) (*httpbody.HttpBody, error) {
	data, err := protojson.Marshal(req)
	if err != nil {
		return nil, err
	}
	return &httpbody.HttpBody{Data: data}, nil
}

type MockInt64QueryService struct{}

func (m *MockInt64QueryService) Int64Query(ctx context.Context, req *Int64QueryRequest) (*httpbody.HttpBody, error) {
	data, err := protojson.Marshal(req)
	if err != nil {
		return nil, err
	}
	return &httpbody.HttpBody{Data: data}, nil
}

type MockUint32QueryService struct{}

func (m *MockUint32QueryService) Uint32Query(ctx context.Context, req *Uint32QueryRequest) (*httpbody.HttpBody, error) {
	data, err := protojson.Marshal(req)
	if err != nil {
		return nil, err
	}
	return &httpbody.HttpBody{Data: data}, nil
}

type MockUint64QueryService struct{}

func (m *MockUint64QueryService) Uint64Query(ctx context.Context, req *Uint64QueryRequest) (*httpbody.HttpBody, error) {
	data, err := protojson.Marshal(req)
	if err != nil {
		return nil, err
	}
	return &httpbody.HttpBody{Data: data}, nil
}

type MockFloatQueryService struct{}

func (m *MockFloatQueryService) FloatQuery(ctx context.Context, req *FloatQueryRequest) (*httpbody.HttpBody, error) {
	data, err := protojson.Marshal(req)
	if err != nil {
		return nil, err
	}
	return &httpbody.HttpBody{Data: data}, nil
}

type MockDoubleQueryService struct{}

func (m *MockDoubleQueryService) DoubleQuery(ctx context.Context, req *DoubleQueryRequest) (*httpbody.HttpBody, error) {
	data, err := protojson.Marshal(req)
	if err != nil {
		return nil, err
	}
	return &httpbody.HttpBody{Data: data}, nil
}

type MockStringQueryService struct{}

func (m *MockStringQueryService) StringQuery(ctx context.Context, req *StringQueryRequest) (*httpbody.HttpBody, error) {
	data, err := protojson.Marshal(req)
	if err != nil {
		return nil, err
	}
	return &httpbody.HttpBody{Data: data}, nil
}

type MockEnumQueryService struct{}

func (m *MockEnumQueryService) EnumQuery(ctx context.Context, req *EnumQueryRequest) (*httpbody.HttpBody, error) {
	data, err := protojson.Marshal(req)
	if err != nil {
		return nil, err
	}
	return &httpbody.HttpBody{Data: data}, nil
}

// ---- Test Cases ----

func TestBoolPath(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router = AppendBoolQueryGonicRoute(router, &MockBoolQueryService{})
	server := httptest.NewServer(router)
	defer server.Close()

	url := server.URL + "/v1/bool?bool=true&opt_bool=false&wrap_bool=true&list_bool=true&list_bool=false&list_wrap_bool=true&list_wrap_bool=false"
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	expected := `{"bool":true, "optBool":false, "wrapBool":true, "listBool":[true, false], "listWrapBool":[true, false]}`
	if strings.ReplaceAll(string(body), " ", "") != strings.ReplaceAll(expected, " ", "") {
		t.Fatalf("body is not equal: got %s, want %s", string(body), expected)
	}
}

func TestInt32Path(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router = AppendInt32QueryGonicRoute(router, &MockInt32QueryService{})
	server := httptest.NewServer(router)
	defer server.Close()

	url := server.URL + "/v1/int32?int32=1&sint32=2&sfixed32=3&opt_int32=4&opt_sint32=5&opt_sfixed32=6&wrap_int32=7&list_int32=1&list_int32=2&list_sint32=1&list_sint32=2&list_sfixed32=1&list_sfixed32=2&list_wrap_int32=1&list_wrap_int32=2"
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	expected := `{"int32":1,"sint32":2,"sfixed32":3,"optInt32":4,"optSint32":5,"optSfixed32":6,"wrapInt32":7,"listInt32":[1,2],"listSint32":[1,2],"listSfixed32":[1,2],"listWrapInt32":[1,2]}`
	if strings.ReplaceAll(string(body), " ", "") != strings.ReplaceAll(expected, " ", "") {
		t.Fatalf("body is not equal: got %s, want %s", string(body), expected)
	}
}

func TestInt64Path(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router = AppendInt64QueryGonicRoute(router, &MockInt64QueryService{})
	server := httptest.NewServer(router)
	defer server.Close()

	url := server.URL + "/v1/int64?int64=10&sint64=20&sfixed64=30&opt_int64=40&opt_sint64=50&opt_sfixed64=60&wrap_int64=70&list_int64=1&list_int64=2&list_sint64=1&list_sint64=2&list_sfixed64=1&list_sfixed64=2&list_wrap_int64=1&list_wrap_int64=2"
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	expected := `{"int64":"10", "sint64":"20", "sfixed64":"30", "optInt64":"40", "optSint64":"50", "optSfixed64":"60", "wrapInt64":"70", "listInt64":["1", "2"], "listSint64":["1", "2"], "listSfixed64":["1", "2"], "listWrapInt64":["1", "2"]}`
	if strings.ReplaceAll(string(body), " ", "") != strings.ReplaceAll(expected, " ", "") {
		t.Fatalf("body is not equal: got %s, want %s", string(body), expected)
	}
}

func TestUint32Path(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router = AppendUint32QueryGonicRoute(router, &MockUint32QueryService{})
	server := httptest.NewServer(router)
	defer server.Close()

	url := server.URL + "/v1/uint32?uint32=1&fixed32=2&opt_uint32=3&opt_fixed32=4&wrap_uint32=5&list_uint32=1&list_uint32=2&list_fixed32=1&list_fixed32=2&list_wrap_uint32=1&list_wrap_uint32=2"
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	expected := `{"uint32":1,"fixed32":2,"optUint32":3,"optFixed32":4,"wrapUint32":5,"listUint32":[1,2],"listFixed32":[1,2],"listWrapUint32":[1,2]}`
	if strings.ReplaceAll(string(body), " ", "") != strings.ReplaceAll(expected, " ", "") {
		t.Fatalf("body is not equal: got %s, want %s", string(body), expected)
	}
}

func TestUint64Path(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router = AppendUint64QueryGonicRoute(router, &MockUint64QueryService{})
	server := httptest.NewServer(router)
	defer server.Close()

	url := server.URL + "/v1/uint64?uint64=10&fixed64=20&opt_uint64=30&opt_fixed64=40&wrap_uint64=50&list_uint64=1&list_uint64=2&list_fixed64=1&list_fixed64=2&list_wrap_uint64=1&list_wrap_uint64=2"
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	expected := `{"uint64":"10", "fixed64":"20", "optUint64":"30", "optFixed64":"40", "wrapUint64":"50", "listUint64":["1", "2"], "listFixed64":["1", "2"], "listWrapUint64":["1", "2"]}`
	if strings.ReplaceAll(string(body), " ", "") != strings.ReplaceAll(expected, " ", "") {
		t.Fatalf("body is not equal: got %s, want %s", string(body), expected)
	}
}

func TestFloatPath(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router = AppendFloatQueryGonicRoute(router, &MockFloatQueryService{})
	server := httptest.NewServer(router)
	defer server.Close()

	url := server.URL + "/v1/float?float=1.23&opt_float=4.56&wrap_float=7.89&list_float=1.23&list_float=3.45&list_wrap_float=4.32&list_wrap_float=5.66"
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	expected := `{"float":1.23, "optFloat":4.56, "wrapFloat":7.89, "listFloat":[1.23, 3.45], "listWrapFloat":[4.32, 5.66]}`
	if strings.ReplaceAll(string(body), " ", "") != strings.ReplaceAll(expected, " ", "") {
		t.Fatalf("body is not equal: got %s, want %s", string(body), expected)
	}
}

func TestDoublePath(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router = AppendDoubleQueryGonicRoute(router, &MockDoubleQueryService{})
	server := httptest.NewServer(router)
	defer server.Close()

	url := server.URL + "/v1/double?double=1.23&opt_double=4.56&wrap_double=7.89&list_double=1.23&list_double=3.45&list_wrap_double=4.32&list_wrap_double=5.66"
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	expected := `{"double":1.23,"optDouble":4.56,"wrapDouble":7.89,"listDouble":[1.23,3.45],"listWrapDouble":[4.32,5.66]}`
	if strings.ReplaceAll(string(body), " ", "") != strings.ReplaceAll(expected, " ", "") {
		t.Fatalf("body is not equal: got %s, want %s", string(body), expected)
	}
}

func TestStringPath(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router = AppendStringQueryGonicRoute(router, &MockStringQueryService{})
	server := httptest.NewServer(router)
	defer server.Close()

	url := server.URL + "/v1/string?string=abc&opt_string=def&wrap_string=ghi&list_string=d3d&list_string=lo-&list_wrap_string=<>d&list_wrap_string={[]}"
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	expected := `{"string":"abc","optString":"def","wrapString":"ghi","listString":["d3d","lo-"],"listWrapString":["<>d","{[]}"]}`
	if strings.ReplaceAll(string(body), " ", "") != strings.ReplaceAll(expected, " ", "") {
		t.Fatalf("body is not equal: got %s, want %s", string(body), expected)
	}
}

func TestEnumPath(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router = AppendEnumQueryGonicRoute(router, &MockEnumQueryService{})
	server := httptest.NewServer(router)
	defer server.Close()

	url := server.URL + "/v1/enum?status=1&opt_status=2&list_status=1&list_status=2"
	resp, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	expected := `{"status":"OK", "optStatus":"CANCELLED", "listStatus":["OK", "CANCELLED"]}`
	if strings.ReplaceAll(string(body), " ", "") != strings.ReplaceAll(expected, " ", "") {
		t.Fatalf("body is not equal: got %s, want %s", string(body), expected)
	}
}
