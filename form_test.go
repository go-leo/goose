package goose

import (
	"errors"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestGetForm(t *testing.T) {
	form := url.Values{}
	form.Set("a", "1")
	getter := func(f url.Values, key string) (int, error) {
		return 42, nil
	}
	v, err := GetForm(nil, form, "a", getter)
	if err != nil || v != 42 {
		t.Errorf("GetForm = %v, %v; want 42, nil", v, err)
	}

	preErr := errors.New("fail")
	v, err = GetForm(preErr, form, "a", getter)
	if err != preErr {
		t.Errorf("GetForm with pre error = %v, %v; want pre error", v, err)
	}
}

func TestGetForm_GenericType(t *testing.T) {
	form := url.Values{}
	form.Set("foo", "bar")
	getter := func(f url.Values, key string) (string, error) {
		return f.Get(key), nil
	}
	v, err := GetForm(nil, form, "foo", getter)
	if err != nil || v != "bar" {
		t.Errorf("GetForm generic = %v, %v; want bar, nil", v, err)
	}
}

func TestGetForm_ErrorPropagation(t *testing.T) {
	form := url.Values{}
	getter := func(f url.Values, key string) (string, error) {
		return "", errors.New("fail")
	}
	v, err := GetForm(nil, form, "foo", getter)
	if err == nil {
		t.Errorf("GetForm should propagate error")
	}
	if v != "" {
		t.Errorf("GetForm should return zero value on error")
	}
}

func TestGetForm_NilForm(t *testing.T) {
	getter := func(f url.Values, key string) (string, error) {
		if f == nil {
			return "ok", nil
		}
		return "", nil
	}
	v, err := GetForm(nil, nil, "foo", getter)
	if err != nil || v != "ok" {
		t.Errorf("GetForm nil form = %v, %v; want ok, nil", v, err)
	}
}

func TestFormGetter_Type(t *testing.T) {
	var _ FormGetter[int]
	var _ FormGetter[string]
}

func TestBreak_Type(t *testing.T) {
	_ = BreakOnError[int]
	_ = BreakOnError[string]
}

func TestFormFromPath(t *testing.T) {
	// 创建一个模拟的HTTP请求，带有路径参数
	req := httptest.NewRequest("GET", "/test", nil)
	// 手动设置路径值（在真实环境中，这通常由路由器处理）
	req.SetPathValue("id", "123")
	req.SetPathValue("name", "test")
	req.SetPathValue("active", "true")

	t.Run("normal case", func(t *testing.T) {
		// 测试正常情况：提取指定的路径参数
		form := FormFromPath(req, "id", "name")

		// 验证返回的form不为nil
		if form == nil {
			t.Error("Expected non-nil url.Values, got nil")
		}
		// 验证提取到正确的值
		if form.Get("id") != "123" {
			t.Errorf("Expected id to be '123', got '%s'", form.Get("id"))
		}
		if form.Get("name") != "test" {
			t.Errorf("Expected name to be 'test', got '%s'", form.Get("name"))
		}
		// 验证未请求的键不存在
		if form.Get("active") != "" {
			t.Errorf("Expected active to be empty, got '%s'", form.Get("active"))
		}
	})

	t.Run("nil keys", func(t *testing.T) {
		// 测试keys为nil的情况
		form := FormFromPath(req)

		// 验证返回nil
		if form != nil {
			t.Error("Expected nil when no keys provided, got non-nil")
		}
	})

	t.Run("empty keys", func(t *testing.T) {
		// 测试空keys切片的情况
		form := FormFromPath(req, []string{}...)

		// 验证返回空的url.Values而不是nil
		if form == nil {
			t.Error("Expected non-nil url.Values, got nil")
		}
		if len(form) != 0 {
			t.Errorf("Expected empty url.Values, got %d values", len(form))
		}
	})

	t.Run("nonexistent key", func(t *testing.T) {
		// 测试请求不存在的键
		form := FormFromPath(req, "nonexistent")

		// 验证返回的form不为nil
		if form == nil {
			t.Error("Expected non-nil url.Values, got nil")
		}
		// 验证不存在的键返回空字符串
		if form.Get("nonexistent") != "" {
			t.Errorf("Expected empty string for nonexistent key, got '%s'", form.Get("nonexistent"))
		}
	})

	t.Run("partial nonexistent keys", func(t *testing.T) {
		// 测试部分键存在，部分不存在的情况
		form := FormFromPath(req, "id", "nonexistent")

		// 验证返回的form不为nil
		if form == nil {
			t.Error("Expected non-nil url.Values, got nil")
		}
		// 验证存在的键有正确值
		if form.Get("id") != "123" {
			t.Errorf("Expected id to be '123', got '%s'", form.Get("id"))
		}
		// 验证不存在的键为空
		if form.Get("nonexistent") != "" {
			t.Errorf("Expected empty string for nonexistent key, got '%s'", form.Get("nonexistent"))
		}
	})
}
