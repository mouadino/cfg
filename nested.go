package cfg

type NestedStructField struct {
	*StructField
}

func (field *NestedStructField) Parse(getter ValueGetter) error {
	err := Parse(
		field.getNestedGetter(getter),
		field.Value.Interface(),
	)
	if err != nil {
		return err
	}
	return nil
}

func (field *NestedStructField) getNestedGetter(getter ValueGetter) ValueGetter {
	return func(keys ...string) interface{} {
		// prepend.
		keys = append([]string{field.Name()}, keys...)
		return getter(keys...)
	}

}
