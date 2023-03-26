package gounit

import (
	"reflect"
	"testing"
)

type T testing.T

func (t *T) AssertInstanceOf(expected, actual interface{}) {
	if reflect.TypeOf(expected) != reflect.TypeOf(actual) {
		t.Helper()
		t.Errorf("Types not equals expected: %d, got: %d", reflect.TypeOf(expected).String(), reflect.TypeOf(actual).String())
	}
}

func (t *T) AssertNotError(err error) {
	if err != nil {
		t.Helper()
		t.Error(err)
	}
}

func (t *T) AssertNil(val interface{}) {
	if reflect.ValueOf(val).IsNil() {
		return
	}
	t.Helper()
	t.Errorf("Value is not nil. got: %s", val)
}

func (t *T) ExpectError(err error) {
	if err == nil {
		t.Helper()
		t.Error("Expected error not present")
	}
}

func (t *T) AssertEqualsInt(expected, actual int) {
	if actual != expected {
		t.Helper()
		t.Errorf("Values not equal expected: %d, got: %d", expected, actual)
	}
}

func (t *T) AssertEqualsString(expected, actual string) {
	if actual != expected {
		t.Helper()
		t.Errorf("Values not equal expected: '%s', got: '%s'", expected, actual)
	}
}

func (t *T) AssertGreaterInt(first, second int) {
	if first <= second {
		t.Helper()
		t.Errorf("%d is not greater than %d", first, second)
	}
}

func (t *T) AssertGreaterOrEqualInt(first, second int) {
	if first < second {
		t.Helper()
		t.Errorf("%d is not greater or equal than %d", first, second)
	}
}

func (t *T) AssertLessInt(first, second int) {
	if first >= second {
		t.Helper()
		t.Errorf("%d is not less than %d", first, second)
	}
}

func (t *T) AssertLessOrEqualInt(first, second int) {
	if first > second {
		t.Helper()
		t.Errorf("%d is not less or equal than %d", first, second)
	}
}

func (t *T) AssertTrue(val bool) {
	if !val {
		t.Helper()
		t.Errorf("Value is not true")
	}
}

func (t *T) AssertFalse(val bool) {
	if val {
		t.Helper()
		t.Errorf("Value is not true")
	}
}

func (t *T) AssertLengthString(expected int, val string) {
	if len(val) != expected {
		t.Helper()
		t.Errorf("Incorect length, expected: %d, got %d", expected, len(val))
	}
}
