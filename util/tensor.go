package util

import (
	ts "github.com/sugarme/gotch/tensor"
	"reflect"
)

// Equal compares 2 tensors in terms of shape, and every element values.
func Equal(tensorA, tensorB *ts.Tensor) bool {
	var equal int64 = 0
	// 1. Compare shape
	if reflect.DeepEqual(tensorA.MustSize(), tensorB.MustSize()) {
		// 2. Compare values
		equal = tensorA.MustEq1(tensorB, false).MustAll(false).Int64Values()[0]
	}
	if equal == 0 {
		return false
	}
	return true
}
