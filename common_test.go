package goose

import (
	"errors"
	"testing"
)

func TestBreak(t *testing.T) {
	// pre error
	pre := errors.New("fail")
	f := BreakOnError[int](pre)
	v, err := f(func() (int, error) { return 1, nil })
	if err != pre || v != 0 {
		t.Errorf("Break with pre error = %v, %v; want 0, pre error", v, err)
	}

	// no error
	f = BreakOnError[int](nil)
	v, err = f(func() (int, error) { return 42, nil })
	if err != nil || v != 42 {
		t.Errorf("Break = %v, %v; want 42, nil", v, err)
	}

	// function error
	f = BreakOnError[int](nil)
	wantErr := errors.New("fail2")
	v, err = f(func() (int, error) { return 0, wantErr })
	if err != wantErr {
		t.Errorf("Break = %v, %v; want 0, wantErr", v, err)
	}
}
