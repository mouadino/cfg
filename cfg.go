package cfg

func Parse(getter ValueGetter, struct_ptr interface{}) error {
	options := Struct{
		Ptr: struct_ptr,
	}
	return options.Parse(getter)
}
