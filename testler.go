package testler

import (
	"reflect"
	"testing"
)

/* Test Helpers */
func Expect(t *testing.T, a interface{}, b interface{}, messages ...string) {
	if a != b {
		var message string = ""
		if len(messages) > 0 {
			message = messages[0] + "\n"
		}
		t.Errorf("%vExpected %v (type %v) - Got %v (type %v)", message, b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

func Refute(t *testing.T, a interface{}, b interface{}, messages ...string) {
	if a == b {
		var message string = ""
		if len(messages) > 0 {
			message = messages[0] + "\n"
		}
		t.Errorf("%vDid not expect %v (type %v) - Got %v (type %v)", message, b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}
