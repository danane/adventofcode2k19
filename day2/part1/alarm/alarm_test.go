package alarm_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/danane/adventofcode2k19/day2/part1/alarm"
)

func TestParseIntcode(t *testing.T) {
	code := `1,9,10,3,2,3,11,0,99,30,40,50`
	got := alarm.ParseIntcode(strings.NewReader(code))
	want := "1,9,10,3,2,3,11,0,99,30,40,50"

	if got != want {
		t.Errorf("got %s, but want %s", got, want)
	}
}

func TestIntcodeToSlice(t *testing.T) {
	code := `1,9,10,3,2,3,11,0,99,30,40,50`

	got := alarm.IntcodeToSlice(code)
	want := []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestExec(t *testing.T) {
	codes := []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}
	want := []int{1, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}
	alarm.Exec(codes, 0)

	if !reflect.DeepEqual(codes, want) {
		t.Errorf("got %v, want %v", codes, want)
	}
}
