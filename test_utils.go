package cfg

import (
	"strings"
)

func emptyGetter(keys ...string) interface{} {
	return nil
}

var values = map[string]interface{}{
	"Str":               "foobar",
	"Int":               1,
	"Nested.Level2_Int": 2,
	"str":               "foobar",
	"int":               1,
	"int64":             int64(1),
	"uint16":            uint16(1),
}

func getter(keys ...string) interface{} {
	return values[strings.Join(keys, ".")]
}
