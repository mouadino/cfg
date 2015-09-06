package cfg

import (
	"reflect"
)

// TODO: transform this to traverse that accept a callable
// that will do the switch type do ...
// TODO: Add support for time.Duration.
// TODO: Get rid of Opt.
type Struct struct {
	Ptr interface{}
}

func (s *Struct) getFields() []IStructField {
	fields := []IStructField{}
	rType := reflect.ValueOf(s.Ptr).Elem()
	for i := 0; i < rType.NumField(); i++ {
		fields = append(fields, NewStructField(rType, i))
	}
	return fields
}

func (s *Struct) Parse(getter ValueGetter) error {
	for _, field := range s.getFields() {
		err := field.Parse(getter)
		if err != nil {
			return err
		}
	}
	return nil
}
