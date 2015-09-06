package cfg

type ValueGetter func(keys ...string) interface{}

type IOption interface {
	Get(ValueGetter) (interface{}, error)
}

type IStructField interface {
	Parse(ValueGetter) error
}
