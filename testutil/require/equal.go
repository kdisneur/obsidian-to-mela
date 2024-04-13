package require

import (
	"reflect"
	"testing"
)

func Equal[T comparable](t *testing.T, want, got T, explanation string) {
	t.Helper()

	if want != got {
		t.Errorf("not equal: %s\nwant: %v\ngot:  %v", explanation, want, got)
		t.FailNow()
	}
}

func DeepEqual(t *testing.T, want, got any, explanation string) {
	t.Helper()

	if !reflect.DeepEqual(want, got) {
		t.Errorf("not equal: %s\nwant: %v\ngot:  %v", explanation, want, got)
		t.FailNow()
	}
}

func NotEmpty[T comparable](t *testing.T, got T, explanation string) {
	t.Helper()

	var empty T

	if got == empty {
		t.Errorf("want empty: %s\ngot: %v", explanation, got)
		t.FailNow()
	}
}

func True(t *testing.T, got bool, explanation string) {
	t.Helper()

	if !got {
		t.Errorf("want false: %s\ngot: %v", explanation, got)
		t.FailNow()
	}
}

func False(t *testing.T, got bool, explanation string) {
	t.Helper()

	if got {
		t.Errorf("want true: %s\ngot: %v", explanation, got)
		t.FailNow()
	}
}

func NoError(t *testing.T, err error, explanation string) {
	t.Helper()

	if err != nil {
		t.Errorf("want no error: %s\ngot: %v", explanation, err)
		t.FailNow()
	}
}
