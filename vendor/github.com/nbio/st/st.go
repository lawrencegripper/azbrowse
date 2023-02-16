// Copyright 2014 nb.io, LLC
// Author: Cameron Walters <cameron@nb.io>

// Package st, pronounced "ghost", is a tiny test framework for
// making short, useful assertions in your Go tests.
//
// To abort a test immediately with t.Fatal, use
// Assert(t, have, want) and Refute(t, have, want)
//
// To allow a test to continue, reporting failure at the end with t.Error, use
// Expect(t, have, want) and Reject(t, have, want)
package st

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

const (
	equal   = "\n%s:%d: should be == \n%s \thave: (%T) %+v\n\twant: (%T) %+v"
	unequal = "\n%s:%d: should be != \n%s \thave: (%T) %+v\n\tand : (%T) %+v"
)

// Errorf is satisfied by testing.T and testing.B.
type Errorf interface {
	Errorf(format string, args ...interface{})
}

// Fatalf is satisfied by testing.T and testing.B.
type Fatalf interface {
	Fatalf(format string, args ...interface{})
}

// Expect calls t.Error and prints a nice comparison message when have != want.
// Especially useful in table-based tests when passing the loop index as iter.
func Expect(t Errorf, have, want interface{}, iter ...int) {
	if !reflect.DeepEqual(have, want) {
		file, line := caller()
		t.Errorf(equal, file, line, exampleNum(iter), have, have, want, want)
	}
}

// Reject calls t.Error and prints a nice comparison message when have == want.
// Especially useful in table-based tests when passing the loop index as iter.
func Reject(t Errorf, have, want interface{}, iter ...int) {
	if reflect.DeepEqual(have, want) {
		file, line := caller()
		t.Errorf(unequal, file, line, exampleNum(iter), have, have, want, want)
	}
}

// Assert calls t.Fatal to abort the test immediately and prints a nice
// comparison message when have != want.
func Assert(t Fatalf, have, want interface{}) {
	if !reflect.DeepEqual(have, want) {
		file, line := caller()
		t.Fatalf(equal, file, line, "", have, have, want, want)
	}
}

// Refute calls t.Fatal to abort the test immediately and prints a nice
// comparison message when have != want.
func Refute(t Fatalf, have, want interface{}) {
	if reflect.DeepEqual(have, want) {
		file, line := caller()
		t.Fatalf(unequal, file, line, "", have, have, want, want)
	}
}

// returns file and line two stack frames above its invocation
func caller() (file string, line int) {
	var ok bool
	_, file, line, ok = runtime.Caller(2)
	if !ok {
		file = "???"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		if slash >= 0 {
			file = file[slash+1:]
		}
	}
	return
}

// returns an example number from the optional zero-based loop iterator n, if
// provided.
func exampleNum(n []int) string {
	if len(n) == 1 {
		return fmt.Sprintf("%d.", n[0])
	}
	return ""
}
