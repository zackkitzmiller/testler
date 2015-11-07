package testler

import (
	"testing"
)

func TestExpectTrue(t *testing.T) {
	Expect(t, true, true)
}

func TestExpectFalse(t *testing.T) {
    Expect(t, false, false)
}

func TestExpectInt(t *testing.T) {
    Expect(t, 1, 1)
}

func TestExpectFloat(t *testing.T) {
    Expect(t, 1.0, 1.0)
}

func TestExpectString(t *testing.T) {
    Expect(t, "string", "string")
}

func TestExpectSuccessWithMessage(t *testing.T) {
    Expect(t, true, true, "true should be true because: reasons")
}

/*
func TestExpectFailWithMessage(t *testing.T) {
    Expect(t, true, false, "true should be true because: reasons")
}
*/

func TestRefuteTrue(t *testing.T) {
    Refute(t, true, false)
}

func TestRefuteFalse(t *testing.T) {
    Refute(t, false, true)
}

func TestRefuteInt(t *testing.T) {
    Refute(t, 1, 2)
}

func TestRefuteFloat(t *testing.T) {
    Refute(t, 1.0, 2.0)
}

func TestRefuteString(t *testing.T) {
    Refute(t, "", "abra")
}

func TestRefuteSuccessWithMessage(t *testing.T) {
    Refute(t, true, false, "true should definitely not be true because: reasons")
}

/*
func TestRefuteFailWithMessage(t *testing.T) {
    Refute(t, true, true, "true should definitely not be true because: reasons")
}
*/
