package cfg

import (
	"fmt"
	"reflect"
	"strconv"
)

func NewStructField(rValue reflect.Value, index int) IStructField {
	field := &StructField{
		Value: rValue.Field(index),
		Type:  rValue.Type().Field(index),
	}
	// TODO: Transform this to switch
	// TODO: Move getOpt switch here.
	if field.Kind() == reflect.Struct {
		return &NestedStructField{field}
	} else if field.Kind() == reflect.Invalid {
		return &InvalidStructField{field}
	} else {
		return field
	}
}

type StructField struct {
	Value reflect.Value
	Type  reflect.StructField
	Tags  reflect.StructTag
}

func (field *StructField) getTag(name string) string {
	return field.Type.Tag.Get(name)
}

func (field *StructField) Kind() reflect.Kind {
	if field.Value.Kind() != reflect.Ptr {
		return field.Value.Kind()
	}
	originalValue := field.Value.Elem()
	if !originalValue.IsValid() {
		return reflect.Invalid // nil pointer.
	}
	return originalValue.Kind()
}

func (field *StructField) Name() string {
	name := field.getTag("name")
	if name == "" {
		name = field.Type.Name
	}
	return name
}

func (field *StructField) Required() (bool, error) {
	var required bool
	var err error

	requiredTag := field.getTag("required")
	if requiredTag != "" {
		required, err = strconv.ParseBool(requiredTag)
	} else {
		required = false
	}
	return required, err
}

func (field *StructField) Default() (interface{}, error) {
	// TODO: Difference between default == "" and not available.
	value := field.getTag("default")
	if value == "" {
		return nil, nil
	}
	return value, nil
}

func (field *StructField) set(newVal interface{}) error {
	if !field.Value.CanSet() {
		return fmt.Errorf("can't set field %s", field.Name())
	}

	// TODO: combine the two switch.
	switch field.Kind() {
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
		err := field.setInt(newVal.(int64))
		if err != nil {
			return err
		}
	case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		err := field.setUInt(newVal.(uint64))
		if err != nil {
			return err
		}
	case reflect.String:
		field.setString(newVal)
	default:
		return fmt.Errorf("can't set field %s: unknown type %s", field.Name(), field.Kind())
	}
	return nil
}

func (field *StructField) setInt(newVal int64) error {
	if !field.Value.OverflowInt(newVal) {
		field.Value.SetInt(newVal)
	} else {
		return fmt.Errorf("int overflow with %t", newVal)
	}
	return nil
}

func (field *StructField) setUInt(newVal uint64) error {
	if !field.Value.OverflowUint(newVal) {
		field.Value.SetUint(newVal)
	} else {
		return fmt.Errorf("uint overflow with %t", newVal)
	}
	return nil
}

func (field *StructField) setString(val interface{}) {
	field.Value.Set(reflect.ValueOf(val))
}

func (field *StructField) getOption() (IOption, error) {
	required, err := field.Required()
	if err != nil {
		return nil, fmt.Errorf("parsing '%s' required tag failed: %s", field.Name(), err)
	}
	defaultValue, err := field.Default()
	if err != nil {
		return nil, fmt.Errorf("parsing '%s' default tag failed: %s", field.Name(), err)
	}

	opt := &Option{
		Name:         field.Name(),
		DefaultValue: defaultValue,
		EnvVar:       field.getTag("env"),
		Required:     required,
	}
	switch field.Kind() {
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
		return &IntOption{opt}, nil
	case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return &UintOption{opt}, nil
	case reflect.String:
		return opt, nil
	default:
		return nil, fmt.Errorf("unhandled type %s", field.Kind())
	}
}

func (field *StructField) Parse(getter ValueGetter) error {
	opt, err := field.getOption()
	if err != nil {
		return err
	}

	optVal, err := opt.Get(getter)
	if err != nil {
		return field.fmtError(err)
	}
	if optVal == nil {
		return nil
	}
	err = field.set(optVal)
	if err != nil {
		return err
	}
	return nil
}

func (field *StructField) fmtError(err error) error {
	// TODO: Field.Name() For nested struct.
	return fmt.Errorf("Field '%s': %s", field.Name(), err)
}
