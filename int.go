package cfg

import (
	"fmt"
	"strconv"
)

type IntOption struct {
	*Option
}

func (opt *IntOption) Get(getter ValueGetter) (interface{}, error) {
	value, err := opt.Option.Get(getter)
	if err != nil {
		return nil, err
	}
	return opt.parseValue(value)
}

// TODO: Refactor parseValue
func (opt *IntOption) parseValue(val interface{}) (interface{}, error) {
	x, err := opt.parseInt(val)
	if err != nil {
		return nil, err
	}
	return int64(x), nil
}

func (opt *IntOption) parseInt(val interface{}) (int64, error) {
	var x64 int64
	var err error
	var ok bool

	// FIXME: Logic is broken.
	switch val.(type) {
	case int64:
		x64, ok = val.(int64)
		if !ok {
			err = fmt.Errorf("can't convert %t to int64", val)
		}
	case int:
		x, ok := val.(int)
		if ok {
			x64 = int64(x)
		}
	case uint16:
		x16, ok := val.(uint64)
		if ok {
			x64 = int64(x16)
		}
	case string:
		x64, err = strconv.ParseInt(val.(string), 10, 64)
	default:
		return 0, fmt.Errorf("unknown integer type %t", val)
	}
	return x64, err
}

type UintOption struct {
	*Option
}

func (opt *UintOption) Get(getter ValueGetter) (interface{}, error) {
	value, err := opt.Option.Get(getter)
	if err != nil {
		return nil, err
	}
	return opt.parseValue(value)
}

func (opt *UintOption) parseValue(val interface{}) (interface{}, error) {
	x, err := opt.parseUint(val)
	if err != nil {
		return nil, err
	}
	return uint64(x), nil
}

func (opt *UintOption) parseUint(val interface{}) (uint64, error) {
	var x64 uint64
	var err error

	switch val.(type) {
	case uint16:
		x16, _ := val.(uint16)
		x64 = uint64(x16)
	// XXX: Reading int and putting it to uint
	// may be dangerous, it's done b/c values from config
	// may come as int directly.
	case int:
		xint, _ := val.(int)
		x64 = uint64(xint)
	case string:
		x64, err = strconv.ParseUint(val.(string), 10, 64)
	default:
		return 0, fmt.Errorf("unknown uint type %t", val)
	}
	return x64, err
}
