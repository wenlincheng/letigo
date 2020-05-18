package util

import (
	"fmt"
	"testing"
)

func TestGenRandNo(t *testing.T) {

	randNo := GenRandNo(1234567)
	fmt.Print(randNo)
}
