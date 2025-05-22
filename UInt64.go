package gqlgen_types

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
)

type UInt64 uint64

func MarshalUInt64(i UInt64) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, fmt.Sprintf("%d", i))
	})
}

func UnmarshalUInt64(v interface{}) (UInt64, error) {
	switch v := v.(type) {
	case string:
		i, err := strconv.Atoi(v)
		if err != nil {
			return 0, fmt.Errorf("string failed to be parsed: %v", err)
		} else {
			return UInt64(i), nil
		}

	case int:
		return UInt64(v), nil
	case int64:
		return UInt64(v), nil
	case json.Number:
		i, err := strconv.Atoi(string(v))
		if err != nil {
			return 0, fmt.Errorf("json.Number failed to be parsed: %v", err)
		} else {
			return UInt64(i), nil
		}
	default:
		return 0, fmt.Errorf("%T is not an int", v)
	}
}
