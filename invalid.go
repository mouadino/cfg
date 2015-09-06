package cfg

import (
	"fmt"
)

type InvalidStructField struct {
	*StructField
}

func (field *InvalidStructField) Parse(_ ValueGetter) error {
	return fmt.Errorf("%v is nil pointer", field.Name())
}
