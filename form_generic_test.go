package goose

import (
	"math"
	"net/url"
	"reflect"
	"testing"
)

func TestGetInt(t *testing.T) {
	form := url.Values{}
	form.Set("a", "123")
	form.Set("b", "-42")
	form.Set("invalid", "abc")

	v, err := GetInt[int64](form, "a")
	if err != nil || v != 123 {
		t.Errorf("GetInt[int64](a) = %v, %v; want 123, nil", v, err)
	}
	v, err = GetInt[int64](form, "b")
	if err != nil || v != -42 {
		t.Errorf("GetInt[int64](b) = %v, %v; want -42, nil", v, err)
	}
	v, err = GetInt[int64](form, "invalid")
	if err == nil {
		t.Errorf("GetInt[int64](invalid) = %v, %v; want error", v, err)
	}
	v, err = GetInt[int64](form, "notfound")
	if err != nil || v != 0 {
		t.Errorf("GetInt[int64](notfound) = %v, %v; want 0, nil", v, err)
	}
}

func TestGetIntPtr(t *testing.T) {
	form := url.Values{}
	form.Set("a", "123")
	ptr, err := GetIntPtr[int64](form, "a")
	if err != nil || ptr == nil || *ptr != 123 {
		t.Errorf("GetIntPtr[int64](a) = %v, %v; want ptr to 123, nil", ptr, err)
	}
	ptr, err = GetIntPtr[int64](form, "notfound")
	if err != nil || ptr == nil || *ptr != 0 {
		t.Errorf("GetIntPtr[int64](notfound) = %v, %v; want ptr to 0, nil", ptr, err)
	}
}

func TestGetIntSlice(t *testing.T) {
	form := url.Values{}
	form["a"] = []string{"1", "2", "3"}
	form["b"] = []string{"x", "2"}
	got, err := GetIntSlice[int64](form, "a")
	want := []int64{1, 2, 3}
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Errorf("GetIntSlice[int64](a) = %v, %v; want %v, nil", got, err, want)
	}
	_, err = GetIntSlice[int64](form, "b")
	if err == nil {
		t.Errorf("GetIntSlice[int64](b) should return error")
	}
	got, err = GetIntSlice[int64](form, "notfound")
	if err != nil || got != nil {
		t.Errorf("GetIntSlice[int64](notfound) = %v, %v; want nil, nil", got, err)
	}
}

func TestGetUint(t *testing.T) {
	form := url.Values{}
	form.Set("a", "123")
	form.Set("invalid", "-1")
	v, err := GetUint[uint64](form, "a")
	if err != nil || v != 123 {
		t.Errorf("GetUint[uint64](a) = %v, %v; want 123, nil", v, err)
	}
	_, err = GetUint[uint64](form, "invalid")
	if err == nil {
		t.Errorf("GetUint[uint64](invalid) should return error")
	}
	v, err = GetUint[uint64](form, "notfound")
	if err != nil || v != 0 {
		t.Errorf("GetUint[uint64](notfound) = %v, %v; want 0, nil", v, err)
	}
}

func TestGetUintPtr(t *testing.T) {
	form := url.Values{}
	form.Set("a", "123")
	ptr, err := GetUintPtr[uint64](form, "a")
	if err != nil || ptr == nil || *ptr != 123 {
		t.Errorf("GetUintPtr[uint64](a) = %v, %v; want ptr to 123, nil", ptr, err)
	}
	ptr, err = GetUintPtr[uint64](form, "notfound")
	if err != nil || ptr == nil || *ptr != 0 {
		t.Errorf("GetUintPtr[uint64](notfound) = %v, %v; want ptr to 0, nil", ptr, err)
	}
}

func TestGetUintSlice(t *testing.T) {
	form := url.Values{}
	form["a"] = []string{"1", "2", "3"}
	form["b"] = []string{"-1", "2"}
	got, err := GetUintSlice[uint64](form, "a")
	want := []uint64{1, 2, 3}
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Errorf("GetUintSlice[uint64](a) = %v, %v; want %v, nil", got, err, want)
	}
	_, err = GetUintSlice[uint64](form, "b")
	if err == nil {
		t.Errorf("GetUintSlice[uint64](b) should return error")
	}
	got, err = GetUintSlice[uint64](form, "notfound")
	if err != nil || got != nil {
		t.Errorf("GetUintSlice[uint64](notfound) = %v, %v; want nil, nil", got, err)
	}
}

func TestGetFloat(t *testing.T) {
	form := url.Values{}
	form.Set("a", "3.14")
	form.Set("invalid", "abc")
	v, err := GetFloat[float64](form, "a")
	if err != nil || math.Abs(v-3.14) > 1e-9 {
		t.Errorf("GetFloat[float64](a) = %v, %v; want 3.14, nil", v, err)
	}
	_, err = GetFloat[float64](form, "invalid")
	if err == nil {
		t.Errorf("GetFloat[float64](invalid) should return error")
	}
	v, err = GetFloat[float64](form, "notfound")
	if err != nil || v != 0 {
		t.Errorf("GetFloat[float64](notfound) = %v, %v; want 0, nil", v, err)
	}
}

func TestGetFloatPtr(t *testing.T) {
	form := url.Values{}
	form.Set("a", "3.14")
	ptr, err := GetFloatPtr[float64](form, "a")
	if err != nil || ptr == nil || math.Abs(*ptr-3.14) > 1e-9 {
		t.Errorf("GetFloatPtr[float64](a) = %v, %v; want ptr to 3.14, nil", ptr, err)
	}
	ptr, err = GetFloatPtr[float64](form, "notfound")
	if err != nil || ptr == nil || *ptr != 0 {
		t.Errorf("GetFloatPtr[float64](notfound) = %v, %v; want ptr to 0, nil", ptr, err)
	}
}

func TestGetFloatSlice(t *testing.T) {
	form := url.Values{}
	form["a"] = []string{"1.1", "2.2", "3.3"}
	form["b"] = []string{"x", "2.2"}
	got, err := GetFloatSlice[float64](form, "a")
	want := []float64{1.1, 2.2, 3.3}
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Errorf("GetFloatSlice[float64](a) = %v, %v; want %v, nil", got, err, want)
	}
	_, err = GetFloatSlice[float64](form, "b")
	if err == nil {
		t.Errorf("GetFloatSlice[float64](b) should return error")
	}
}
