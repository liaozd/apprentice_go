package data

import (
	"fmt"
	"strconv"
)

// Declare a custom Runtime type, which has the underlying type int32 (the same as our
// Movie struct field).
type Runtime int32

// We’re deliberately using a value receiver for our MarshalJSON() method rather than a
// pointerreceiverlikefunc (r *Runtime) MarshalJSON().Thisgivesusmoreflexibility because
// it means that our custom JSON encoding will work on both Runtime values and pointers
// to Runtime values. As Effective Go mentions:
//
// 		The rule about pointers vs. values for receivers is that value methods can be invoked on pointers and values,
//		but pointer methods can only be invoked on pointers.
//
// Implement a MarshalJSON() method on the Runtime type so that it satisfies the
// json.Marshaler interface. T his should return the JSON-encoded value for the movie
// runtime (in our case, it will return a string in the format "<runtime> mins").
func (r Runtime) MarshalJSON() ([]byte, error) {
	// Generate a string containing the movie runtime in the required format.
	jsonValue := fmt.Sprintf("%d mins", r)

	// strconv.Quote() function on the string to wrap it in double quotes. It
	// needs to surrounded by double quotes in order to be
	quotedJSONValue := strconv.Quote(jsonValue)

	return []byte(quotedJSONValue), nil
}
