// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package constant

import (
	shimjson "github.com/SignifyHQ/rain-sdk-go/internal/encoding/json"
)

type Constant[T any] interface {
	Default() T
}

// ValueOf gives the default value of a constant from its type. It's helpful when
// constructing constants as variants in a one-of. Note that empty structs are
// marshalled by default. Usage: constant.ValueOf[constant.Foo]()
func ValueOf[T Constant[T]]() T {
	var t T
	return t.Default()
}

type Collateral string // Always "collateral"
type Fee string        // Always "fee"
type Payment string    // Always "payment"
type Spend string      // Always "spend"

func (c Collateral) Default() Collateral { return "collateral" }
func (c Fee) Default() Fee               { return "fee" }
func (c Payment) Default() Payment       { return "payment" }
func (c Spend) Default() Spend           { return "spend" }

func (c Collateral) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c Fee) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c Payment) MarshalJSON() ([]byte, error)    { return marshalString(c) }
func (c Spend) MarshalJSON() ([]byte, error)      { return marshalString(c) }

type constant[T any] interface {
	Constant[T]
	*T
}

func marshalString[T ~string, PT constant[T]](v T) ([]byte, error) {
	var zero T
	if v == zero {
		v = PT(&v).Default()
	}
	return shimjson.Marshal(string(v))
}
