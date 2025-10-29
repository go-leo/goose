package goose

import (
	"net/url"
	"reflect"
	"testing"

	"google.golang.org/protobuf/types/known/wrapperspb"
)


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
