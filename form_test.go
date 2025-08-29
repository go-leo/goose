package goose

import (
	"math"
	"net/url"
	"reflect"
	"testing"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestGetInt32(t *testing.T) {
	form := url.Values{}
	form.Set("a", "123")
	v, err := GetInt32(form, "a")
	if err != nil || v != 123 {
		t.Errorf("GetInt32(a) = %v, %v; want 123, nil", v, err)
	}
}

func TestGetInt32Ptr(t *testing.T) {
	form := url.Values{}
	form.Set("a", "123")
	ptr, err := GetInt32Ptr(form, "a")
	if err != nil || ptr == nil || *ptr != 123 {
		t.Errorf("GetInt32Ptr(a) = %v, %v; want ptr to 123, nil", ptr, err)
	}
}

func TestGetInt64(t *testing.T) {
	form := url.Values{}
	form.Set("a", "123456789")
	v, err := GetInt64(form, "a")
	if err != nil || v != 123456789 {
		t.Errorf("GetInt64(a) = %v, %v; want 123456789, nil", v, err)
	}
}

func TestGetInt64Ptr(t *testing.T) {
	form := url.Values{}
	form.Set("a", "123456789")
	ptr, err := GetInt64Ptr(form, "a")
	if err != nil || ptr == nil || *ptr != 123456789 {
		t.Errorf("GetInt64Ptr(a) = %v, %v; want ptr to 123456789, nil", ptr, err)
	}
}

func TestGetUint32(t *testing.T) {
	form := url.Values{}
	form.Set("a", "123")
	v, err := GetUint32(form, "a")
	if err != nil || v != 123 {
		t.Errorf("GetUint32(a) = %v, %v; want 123, nil", v, err)
	}
}

func TestGetUint32Ptr(t *testing.T) {
	form := url.Values{}
	form.Set("a", "123")
	ptr, err := GetUint32Ptr(form, "a")
	if err != nil || ptr == nil || *ptr != 123 {
		t.Errorf("GetUint32Ptr(a) = %v, %v; want ptr to 123, nil", ptr, err)
	}
}

func TestGetUint64(t *testing.T) {
	form := url.Values{}
	form.Set("a", "123456789")
	v, err := GetUint64(form, "a")
	if err != nil || v != 123456789 {
		t.Errorf("GetUint64(a) = %v, %v; want 123456789, nil", v, err)
	}
}

func TestGetUint64Ptr(t *testing.T) {
	form := url.Values{}
	form.Set("a", "123456789")
	ptr, err := GetUint64Ptr(form, "a")
	if err != nil || ptr == nil || *ptr != 123456789 {
		t.Errorf("GetUint64Ptr(a) = %v, %v; want ptr to 123456789, nil", ptr, err)
	}
}

func TestGetFloat32(t *testing.T) {
	form := url.Values{}
	form.Set("a", "3.14")
	v, err := GetFloat32(form, "a")
	if err != nil || math.Abs(float64(v-3.14)) > 1e-6 {
		t.Errorf("GetFloat32(a) = %v, %v; want 3.14, nil", v, err)
	}
}

func TestGetFloat32Ptr(t *testing.T) {
	form := url.Values{}
	form.Set("a", "3.14")
	ptr, err := GetFloat32Ptr(form, "a")
	if err != nil || ptr == nil || math.Abs(float64(*ptr-3.14)) > 1e-6 {
		t.Errorf("GetFloat32Ptr(a) = %v, %v; want ptr to 3.14, nil", ptr, err)
	}
}

func TestGetFloat32Slice(t *testing.T) {
	form := url.Values{}
	form["a"] = []string{"1.1", "2.2"}
	got, err := GetFloat32Slice(form, "a")
	want := []float32{1.1, 2.2}
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Errorf("GetFloat32Slice(a) = %v, %v; want %v, nil", got, err, want)
	}
}

func TestGetFloat64(t *testing.T) {
	form := url.Values{}
	form.Set("a", "3.1415")
	v, err := GetFloat64(form, "a")
	if err != nil || math.Abs(v-3.1415) > 1e-9 {
		t.Errorf("GetFloat64(a) = %v, %v; want 3.1415, nil", v, err)
	}
}

func TestGetFloat64Ptr(t *testing.T) {
	form := url.Values{}
	form.Set("a", "3.1415")
	ptr, err := GetFloat64Ptr(form, "a")
	if err != nil || ptr == nil || math.Abs(*ptr-3.1415) > 1e-9 {
		t.Errorf("GetFloat64Ptr(a) = %v, %v; want ptr to 3.1415, nil", ptr, err)
	}
}

func TestGetFloat64Slice(t *testing.T) {
	form := url.Values{}
	form["a"] = []string{"1.1", "2.2"}
	got, err := GetFloat64Slice(form, "a")
	want := []float64{1.1, 2.2}
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Errorf("GetFloat64Slice(a) = %v, %v; want %v, nil", got, err, want)
	}
}

func TestGetBool(t *testing.T) {
	form := url.Values{}
	form.Set("a", "true")
	v, err := GetBool(form, "a")
	if err != nil || v != true {
		t.Errorf("GetBool(a) = %v, %v; want true, nil", v, err)
	}
	v, err = GetBool(form, "notfound")
	if err != nil || v != false {
		t.Errorf("GetBool(notfound) = %v, %v; want false, nil", v, err)
	}
}

func TestGetBoolPtr(t *testing.T) {
	form := url.Values{}
	form.Set("a", "true")
	ptr, err := GetBoolPtr(form, "a")
	if err != nil || ptr == nil || *ptr != true {
		t.Errorf("GetBoolPtr(a) = %v, %v; want ptr to true, nil", ptr, err)
	}
}

func TestGetBoolSlice(t *testing.T) {
	form := url.Values{}
	form["a"] = []string{"true", "false"}
	got, err := GetBoolSlice(form, "a")
	want := []bool{true, false}
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Errorf("GetBoolSlice(a) = %v, %v; want %v, nil", got, err, want)
	}
	got, err = GetBoolSlice(form, "notfound")
	if err != nil || got != nil {
		t.Errorf("GetBoolSlice(notfound) = %v, %v; want nil, nil", got, err)
	}
}

func TestGetBoolValue(t *testing.T) {
	form := url.Values{}
	form.Set("a", "true")
	got, err := GetBoolValue(form, "a")
	want := wrapperspb.Bool(true)
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Errorf("GetBoolValue(a) = %v, %v; want %v, nil", got, err, want)
	}
}

func TestGetBoolValueSlice(t *testing.T) {
	form := url.Values{}
	form["a"] = []string{"true", "false"}
	got, err := GetBoolValueSlice(form, "a")
	want := []*wrapperspb.BoolValue{wrapperspb.Bool(true), wrapperspb.Bool(false)}
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Errorf("GetBoolValueSlice(a) = %v, %v; want %v, nil", got, err, want)
	}
}

func TestGetInt32Value(t *testing.T) {
	form := url.Values{}
	form.Set("a", "123")
	got, err := GetInt32Value(form, "a")
	want := wrapperspb.Int32(123)
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Errorf("GetInt32Value(a) = %v, %v; want %v, nil", got, err, want)
	}
}

func TestGetInt32ValueSlice(t *testing.T) {
	form := url.Values{}
	form["a"] = []string{"1", "2"}
	got, err := GetInt32ValueSlice(form, "a")
	want := []*wrapperspb.Int32Value{wrapperspb.Int32(1), wrapperspb.Int32(2)}
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Errorf("GetInt32ValueSlice(a) = %v, %v; want %v, nil", got, err, want)
	}
}

func TestGetInt64Value(t *testing.T) {
	form := url.Values{}
	form.Set("a", "123456789")
	got, err := GetInt64Value(form, "a")
	want := wrapperspb.Int64(123456789)
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Errorf("GetInt64Value(a) = %v, %v; want %v, nil", got, err, want)
	}
}

func TestGetInt64ValueSlice(t *testing.T) {
	form := url.Values{}
	form["a"] = []string{"1", "2"}
	got, err := GetInt64ValueSlice(form, "a")
	want := []*wrapperspb.Int64Value{wrapperspb.Int64(1), wrapperspb.Int64(2)}
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Errorf("GetInt64ValueSlice(a) = %v, %v; want %v, nil", got, err, want)
	}
}

func TestGetUint32Value(t *testing.T) {
	form := url.Values{}
	form.Set("a", "123")
	got, err := GetUint32Value(form, "a")
	want := wrapperspb.UInt32(123)
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Errorf("GetUint32Value(a) = %v, %v; want %v, nil", got, err, want)
	}
}

func TestGetUint32ValueSlice(t *testing.T) {
	form := url.Values{}
	form["a"] = []string{"1", "2"}
	got, err := GetUint32ValueSlice(form, "a")
	want := []*wrapperspb.UInt32Value{wrapperspb.UInt32(1), wrapperspb.UInt32(2)}
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Errorf("GetUint32ValueSlice(a) = %v, %v; want %v, nil", got, err, want)
	}
}

func TestGetUint64Value(t *testing.T) {
	form := url.Values{}
	form.Set("a", "123456789")
	got, err := GetUint64Value(form, "a")
	want := wrapperspb.UInt64(123456789)
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Errorf("GetUint64Value(a) = %v, %v; want %v, nil", got, err, want)
	}
}

func TestGetUint64ValueSlice(t *testing.T) {
	form := url.Values{}
	form["a"] = []string{"1", "2"}
	got, err := GetUint64ValueSlice(form, "a")
	want := []*wrapperspb.UInt64Value{wrapperspb.UInt64(1), wrapperspb.UInt64(2)}
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Errorf("GetUint64ValueSlice(a) = %v, %v; want %v, nil", got, err, want)
	}
}

func TestGetFloat32Value(t *testing.T) {
	form := url.Values{}
	form.Set("a", "3.14")
	got, err := GetFloat32Value(form, "a")
	want := wrapperspb.Float(3.14)
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Errorf("GetFloat32Value(a) = %v, %v; want %v, nil", got, err, want)
	}
}

func TestGetFloat32ValueSlice(t *testing.T) {
	form := url.Values{}
	form["a"] = []string{"1.1", "2.2"}
	got, err := GetFloat32ValueSlice(form, "a")
	want := []*wrapperspb.FloatValue{wrapperspb.Float(1.1), wrapperspb.Float(2.2)}
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Errorf("GetFloat32ValueSlice(a) = %v, %v; want %v, nil", got, err, want)
	}
}

func TestGetFloat64Value(t *testing.T) {
	form := url.Values{}
	form.Set("a", "3.1415")
	got, err := GetFloat64Value(form, "a")
	want := wrapperspb.Double(3.1415)
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Errorf("GetFloat64Value(a) = %v, %v; want %v, nil", got, err, want)
	}
}

func TestGetFloat64ValueSlice(t *testing.T) {
	form := url.Values{}
	form["a"] = []string{"1.1", "2.2"}
	got, err := GetFloat64ValueSlice(form, "a")
	want := []*wrapperspb.DoubleValue{wrapperspb.Double(1.1), wrapperspb.Double(2.2)}
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Errorf("GetFloat64ValueSlice(a) = %v, %v; want %v, nil", got, err, want)
	}
}

