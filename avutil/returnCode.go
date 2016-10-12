// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Corbatto (luca@corbatto.de)

package avutil

import "fmt"

// ReturnCode is a wrapper for return codes.
// It implements the error interface but is also capable of giving you the actual return code.
type ReturnCode struct {
	code int
}

// NewReturnCode creates a new ReturnCode instance
func NewReturnCode(code int) ReturnCode {
	return ReturnCode{code}
}

// Error is the implementation of the error interface.
// It returns a descriptive string of the return code.
//
// Warning: It panics if the returned code is positive (or 0), because those values are no errors. Please use ReturnCode.Ok() to check if it's an error or not.
func (r ReturnCode) Error() string {
	if r.Ok() {
		panic("This is not an error. Please use ReturnCode.Ok() to check that.")
	}
	return fmt.Sprintf("%s (%d)", Strerror(r.code), r.code)
}

// Ok returns true if the return code does not resemble an error.
func (r ReturnCode) Ok() bool {
	return r.code >= 0
}

// Code returns the original return code.
// This can be used e.g. if a positive return code resembles the size of something.
//
// This can also be used for comparison against the AVERROR_xxx values.
func (r ReturnCode) Code() int {
	return r.code
}

// IsOneOf is a convenience method that checks if the return codes equals one of the given ones.
//
// Usage:
//     func YourFunction() error {
//         code := avxxx.SomeFunction()
//         if !code.Ok() && !code.IsOneOf(avutil.AVERROR_EAGAIN, avutil.AVERROR_EOF) {
//             return code
//         }
//     }
func (r ReturnCode) IsOneOf(codes ...int) bool {
	for _, c := range codes {
		if c == r.code {
			return true
		}
	}
	return false
}
