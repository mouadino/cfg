package cfg

import (
	"errors"
	"fmt"
	"os"
)

var (
	RequiredError = errors.New("required value")
)

type Option struct {
	Name         string
	EnvVar       string
	Required     bool
	DefaultValue interface{}
}

func (opt *Option) Get(getter ValueGetter) (interface{}, error) {
	return opt.getValue(getter)
}

func (opt *Option) getValue(getter ValueGetter) (interface{}, error) {
	value := getter(opt.Name)
	if value == nil && opt.EnvVar != "" {
		value = os.Getenv(opt.EnvVar)
	}
	if value == nil && opt.DefaultValue != nil {
		value = opt.DefaultValue
	}
	if value == nil && opt.Required == true {
		return nil, RequiredError
	}
	return value, nil
}

func (opt *Option) String() string {
	return fmt.Sprintf("Option '%v' (default: %v, required: %v", opt.Name, opt.DefaultValue, opt.Required)
}
