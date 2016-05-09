package assert

import (
	"fmt"
	"testing"
)

// Use ...
func Use(t *testing.T) func(interface{}) *matcher {
	return func(expect interface{}) *matcher {
		return &matcher{expect, t}
	}
}

type matcher struct {
	actual interface{}
	t      *testing.T
}

func (m *matcher) Is(expect interface{}) *matcher {
	switch actual := m.actual.(type) {

	case int:
		expect := expect.(int)

		if expect != actual {
			m.t.Fatal(&assertFailInt{actual: actual, expect: expect})
		}
	case string:
		expect := expect.(string)

		if expect != actual {
			m.t.Fatal(&assertFailString{actual: actual, expect: expect})
		}
	case bool:
		expect := expect.(bool)

		if expect != actual {
			m.t.Fatal(&assertFailBool{actual: actual, expect: expect})
		}
	case []string:
		expect := expect.([]string)

		if err := stringArrayEquals(expect, actual); err != nil {
			m.t.Fatal(err)
		}

	default:
		m.t.Fatalf("Assert Not Implements! %s", expect)
	}
	return m
}

func (m *matcher) IsNotNil() *matcher {
	if m.actual == nil {
		m.t.Fatalf("expect not nil but nil")
	}
	return m
}

type assertFail interface {
	Error() string
}

func stringArrayEquals(expect, actual []string) assertFail {
	if len(expect) != len(actual) {
		return &assertFailDifferentLength{
			expect: len(expect),
			actual: len(actual),
		}
	}

	for i := 0; i < len(expect); i++ {
		if expect[i] != actual[i] {
			return &assertFailNotEqualItem{
				expect: expect[i],
				actual: actual[i],
				index:  i,
			}
		}
	}
	return nil
}

type assertFailDifferentLength struct {
	expect, actual int
}

func (fail *assertFailDifferentLength) Error() string {
	return fmt.Sprintf("different length : expect %d, but %d", fail.expect, fail.actual)
}

type assertFailNotEqualItem struct {
	expect, actual string
	index          int
}

func (fail *assertFailNotEqualItem) Error() string {
	return fmt.Sprintf("different item at %d : expect %s, but %s", fail.index, fail.expect, fail.actual)
}

type assertFailInt struct {
	expect, actual int
}

func (fail *assertFailInt) Error() string {
	return fmt.Sprintf("expect %d, but %d", fail.expect, fail.actual)
}

type assertFailString struct {
	expect, actual string
}

func (fail *assertFailString) Error() string {
	return fmt.Sprintf("expect %s, but %s", fail.expect, fail.actual)
}

type assertFailBool struct {
	expect, actual bool
}

func (fail *assertFailBool) Error() string {
	return fmt.Sprintf("expect %t, but %t", fail.expect, fail.actual)
}
