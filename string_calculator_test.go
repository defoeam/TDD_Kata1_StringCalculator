package string_calculator

import (
	"testing"
)

func TestAdd_BaseCase(t *testing.T) {
	got, _ := newStringCalculator().add("")
	want := 0

	if got != want {
		t.Errorf("Add() = %d; want %d", got, want)
	}
}

func TestAdd_SingleArg(t *testing.T) {
	got, _ := newStringCalculator().add("1")
	want := 1

	if got != want {
		t.Errorf("Add() = %d; want %d", got, want)
	}
}

func TestAdd_DualArg(t *testing.T) {
	got, _ := newStringCalculator().add("1,2")
	want := 3

	if got != want {
		t.Errorf("Add() = %d; want %d", got, want)
	}
}

func TestAdd_NArgs(t *testing.T) {
	got, _ := newStringCalculator().add("1,2,3,4,5")
	want := 15

	if got != want {
		t.Errorf("Add() = %d; want %d", got, want)
	}
}

func TestAdd_NewLineDelimiter(t *testing.T) {
	got, _ := newStringCalculator().add("1\n2,3")
	want := 6

	if got != want {
		t.Errorf("Add() = %d; want %d", got, want)
	}
}

func TestAdd_CustomDelimiters(t *testing.T) {
	// “//[delimiter]\n[numbers…]”
	// Example: “//;\n1;2” == 3
	got, _ := newStringCalculator().add("//;\n1;2")
	want := 3

	if got != want {
		t.Errorf("Add() = %d; want %d", got, want)
	}
}

func TestAdd_ErrorOnNegative(t *testing.T) {
	_, err := newStringCalculator().add("-1")
	want := "negative numbers are not allowed: -1"

	if err.Error() != want {
		t.Errorf("Add() did not return correct error for negative input; want %s", want)
	}
}

func TestAdd_MultipleNegative(t *testing.T) {
	_, err := newStringCalculator().add("-1,-2")
	want := "negative numbers are not allowed: -1, -2"

	if err.Error() != want {
		t.Errorf("Add() did not return correct error for multiple negative input; want %s", want)
	}
}

func TestGetCalledCount(t *testing.T) {
	// Setup
	sc := newStringCalculator()
	sc.add("1,2,3") // Call the function to increment the count
	sc.add("4,5,6") // Call it again

	got := sc.getCalledCount()
	want := 2

	if got != want {
		t.Errorf("GetCalledCount() = %d; want %d", got, want)
	}
}

func TestAdd_IgnoreGreaterThanOneThousand(t *testing.T) {
	got, _ := newStringCalculator().add("1001,2")
	want := 2 // 1001 is ignored

	if got != want {
		t.Errorf("Add() = %d; want %d", got, want)
	}
}

func TestAdd_AnyLengthCustomDelimiter(t *testing.T) {
	// Custom delimiter of any length
	got, _ := newStringCalculator().add("//[***]\n1***2***3")
	want := 6

	if got != want {
		t.Errorf("Add() = %d; want %d", got, want)
	}
}

func TestAdd_AllowSeveralSingleLengthDelimiters(t *testing.T) {
	// Allow multiple delimiters
	got, _ := newStringCalculator().add("//[;][*]\n1;2*3")
	want := 6

	if got != want {
		t.Errorf("Add() = %d; want %d", got, want)
	}
}

func TestAdd_AllowSeveralMultiLengthDelimiters(t *testing.T) {
	// Allow multiple delimiters of any length
	got, _ := newStringCalculator().add("//[;;][**]\n1;;2**3")
	want := 6

	if got != want {
		t.Errorf("Add() = %d; want %d", got, want)
	}
}
