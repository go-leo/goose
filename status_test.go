package goose

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

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
