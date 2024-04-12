package require

import "testing"

func Equal[T comparable](t *testing.T, want, got T, explanation string) {
	t.Helper()

	if want != got {
		t.Errorf("not equal: %s\nwant: %v\ngot:  %v")
		t.FailNow()
	}
}

func NotEmpty[T comparable](t *testing.T, got T, explanation string) {
	t.Helper()

	var empty T

	if got == empty {
		t.Errorf("want empty: %s\ngot: %v")
		t.FailNow()
	}
}

func True(t *testing.T, got bool, explanation string) {
	t.Helper()

	if !got {
		t.Errorf("want false: %s\ngot: %v")
		t.FailNow()
	}
}

func False(t *testing.T, got bool, explanation string) {
	t.Helper()

	if got {
		t.Errorf("want true: %s\ngot: %v")
		t.FailNow()
	}
}
