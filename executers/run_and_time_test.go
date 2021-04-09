package executers

import (
	"fmt"
	"testing"
)

func HelloTest(s string) []string {
	if s == "" {
		fmt.Println("NOTHING TO SHOW HERE")
	}

	fmt.Println("JUST BEFORE RETURN")
	fmt.Println(s)
	return []string{s, "another", "test"}
}

func TestTester(t *testing.T) {
	fmt.Println("TestTester")

	fro := TesterExecutor.RunAndTime(HelloTest, "Hello My Test")

	fmt.Println("FRO ->", fro)
}
