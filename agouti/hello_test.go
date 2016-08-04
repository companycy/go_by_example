package main

import (
	"testing"  // 2
)

// 3.2
type sayTwiceTest struct {
	input          string
	expectedOutput string
}

// 3.3
var sayTwiceTests = []sayTwiceTest{
	{"Hello", "HelloHello"},
	{"Bye", "ByeBye"},
}

func TestSayTwice(t *testing.T) {
	for _, test := range sayTwiceTests {  // 3.4
		if output := SayTwice(test.input); output != test.expectedOutput {
			t.Errorf("Expected SayTwice to return (%s), Found (%s)\n", test.expectedOutput, output)
		}
	}
}
