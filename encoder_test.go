package gonic

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"google.golang.org/genproto/googleapis/api/httpbody"
	rpchttp "google.golang.org/genproto/googleapis/rpc/http"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// --- Mocks ---

type mockProto struct {
	proto.Message
}

func (m *mockProto) Reset()         {}
func (m *mockProto) String() string { return "mock" }
func (m *mockProto) ProtoMessage()  {}

type jsonErr struct{}

func (jsonErr) Error() string { return "json error" }
func (jsonErr) MarshalJSON() ([]byte, error) {
	return []byte(`{"msg":"json error"}`), nil
}

type headerErr struct{}

func (headerErr) Error() string { return "header error" }
func (headerErr) Headers() http.Header {
	h := http.Header{}
	h.Set("X-Test", "1")
	return h
}

type statusErr struct{}

func (statusErr) Error() string { return "status error" }
func (statusErr) StatusCode() int {
	return 418
}

// --- Tests ---

func TestDefaultTransformResponse(t *testing.T) {
	ctx := context.Background()
	msg := &mockProto{}
	got := DefaultTransformResponse(ctx, msg)
	if got != msg {
		t.Errorf("DefaultTransformResponse should return input message")
	}
}

func TestDefaultEncodeError_plain(t *testing.T) {
	rr := httptest.NewRecorder()
	DefaultEncodeError(context.Background(), errors.New("plain error"), rr)
	resp := rr.Result()
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("status = %d, want %d", resp.StatusCode, http.StatusInternalServerError)
	}
	if ct := resp.Header.Get(ContentTypeKey); ct != "text/plain; charset=utf-8" {
		t.Errorf("Content-Type = %q, want text/plain", ct)
	}
	if !bytes.Contains(body, []byte("plain error")) {
		t.Errorf("body = %q, want contains 'plain error'", body)
	}
}

func TestDefaultEncodeError_json(t *testing.T) {
	rr := httptest.NewRecorder()
	DefaultEncodeError(context.Background(), jsonErr{}, rr)
	resp := rr.Result()
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if ct := resp.Header.Get(ContentTypeKey); ct != JsonContentType {
		t.Errorf("Content-Type = %q, want %q", ct, JsonContentType)
	}
	if !bytes.Contains(body, []byte(`"msg":"json error"`)) {
		t.Errorf("body = %q, want contains json error", body)
	}
}

func TestDefaultEncodeError_header(t *testing.T) {
	rr := httptest.NewRecorder()
	DefaultEncodeError(context.Background(), headerErr{}, rr)
	resp := rr.Result()
	defer resp.Body.Close()
	if resp.Header.Get("X-Test") != "1" {
		t.Errorf("X-Test header not set")
	}
}

func TestDefaultEncodeError_status(t *testing.T) {
	rr := httptest.NewRecorder()
	DefaultEncodeError(context.Background(), statusErr{}, rr)
	resp := rr.Result()
	defer resp.Body.Close()
	if resp.StatusCode != 418 {
		t.Errorf("status = %d, want 418", resp.StatusCode)
	}
}

func TestEncodeResponse(t *testing.T) {
	rr := httptest.NewRecorder()
	msg := &httpbody.HttpBody{
		ContentType: "application/test",
		Data:        []byte("hello"),
	}
	opts := protojson.MarshalOptions{}
	err := EncodeResponse(context.Background(), rr, msg, opts)
	if err != nil {
		t.Fatalf("EncodeResponse error: %v", err)
	}
	resp := rr.Result()
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("status = %d, want 200", resp.StatusCode)
	}
	if ct := resp.Header.Get(ContentTypeKey); ct != JsonContentType {
		t.Errorf("Content-Type = %q, want %q", ct, JsonContentType)
	}
}

func TestEncodeHttpBody(t *testing.T) {
	rr := httptest.NewRecorder()
	msg := &httpbody.HttpBody{
		ContentType: "application/test",
		Data:        []byte("hello"),
	}
	err := EncodeHttpBody(context.Background(), rr, msg)
	if err != nil {
		t.Fatalf("EncodeHttpBody error: %v", err)
	}
	resp := rr.Result()
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("status = %d, want 200", resp.StatusCode)
	}
	if ct := resp.Header.Get(ContentTypeKey); ct != "application/test" {
		t.Errorf("Content-Type = %q, want application/test", ct)
	}
	body, _ := io.ReadAll(resp.Body)
	if !bytes.Equal(body, []byte("hello")) {
		t.Errorf("body = %q, want %q", body, "hello")
	}
}

func TestEncodeHttpResponse(t *testing.T) {
	rr := httptest.NewRecorder()
	msg := &rpchttp.HttpResponse{
		Status: 201,
		Body:   []byte("abc"),
		Headers: []*rpchttp.HttpHeader{
			{Key: "X-Foo", Value: "bar"},
			{Key: "X-Foo", Value: "baz"},
		},
	}
	err := EncodeHttpResponse(context.Background(), rr, msg)
	if err != nil {
		t.Fatalf("EncodeHttpResponse error: %v", err)
	}
	resp := rr.Result()
	defer resp.Body.Close()
	if resp.StatusCode != 201 {
		t.Errorf("status = %d, want 201", resp.StatusCode)
	}
	if got := resp.Header["X-Foo"]; !reflect.DeepEqual(got, []string{"bar", "baz"}) {
		t.Errorf("X-Foo header = %v, want [bar baz]", got)
	}
	body, _ := io.ReadAll(resp.Body)
	if !bytes.Equal(body, []byte("abc")) {
		t.Errorf("body = %q, want %q", body, "abc")
	}
}
