package cfg

import (
	"os"
	"testing"
)

type optionsDefaultExample struct {
	Str  string `default:"127.0.0.1"`
	Int  int    `default:"4080"`
	Uint uint16 `default:"16"`
}

func TestDefaultsOpt(t *testing.T) {
	opt := &optionsDefaultExample{}
	err := Parse(emptyGetter, opt)
	if err != nil {
		t.Fatalf("Parse failed with: '%s'", err)
	}
	if opt.Str != "127.0.0.1" {
		t.Fatalf("Parse struct Str failed, got: '%s' want: '127.0.0.1'", opt.Str)
	}
	if opt.Int != 4080 {
		t.Fatalf("Parse struct Int failed, got: '%s' want: 4080", opt.Int)
	}
	if opt.Uint != 16 {
		t.Fatalf("Parse struct Int failed, got: '%s' want: 16", opt.Uint)
	}
}

type optionsEnvExample struct {
	Str  string `env:"CFG_STR"`
	Int  int    `env:"CFG_INT"`
	Uint uint16 `env:"CFG_UINT"`
}

func TestEnvOpt(t *testing.T) {
	os.Setenv("CFG_STR", "127.0.0.1")
	os.Setenv("CFG_INT", "4080")
	os.Setenv("CFG_UINT", "16")

	opt := &optionsEnvExample{}
	err := Parse(emptyGetter, opt)
	if err != nil {
		t.Fatalf("Parse failed with: '%s'", err)
	}
	if opt.Str != "127.0.0.1" {
		t.Fatalf("Parse struct Str failed, got: '%s' want: '127.0.0.1'", opt.Str)
	}
	if opt.Int != 4080 {
		t.Fatalf("Parse struct Int failed, got: '%s' want: 4080", opt.Int)
	}
	if opt.Uint != 16 {
		t.Fatalf("Parse struct Int failed, got: '%s' want: 16", opt.Uint)
	}
}

type optionsRequiredExample struct {
	Str string `required:"true"`
	Int int
}

func TestRequiredOpt(t *testing.T) {
	opt := &optionsRequiredExample{}
	err := Parse(emptyGetter, opt)
	if err == nil {
		t.Fatalf("Parsing missing required field didn't fail, got %v", opt)
	}
	msg := "Field 'Str': required value"
	if err.Error() != msg {
		t.Fatalf("Error message didn't match, got: %v want: %v", err.Error(), msg)
	}
}

type level2Options struct {
	Level2_Int int `default:"4080"`
}

type nestedOptionsExample struct {
	Nested *level2Options
	Str    string `default:"127.0.0.1"`
}

func TestNestedOpt(t *testing.T) {
	opt := &nestedOptionsExample{Nested: &level2Options{}}
	err := Parse(getter, opt)
	if err != nil {
		t.Fatalf("Parse failed with: '%s'", err)
	}
	if opt.Str != "foobar" {
		t.Fatalf("Parse struct Str failed, got: '%s' want: 'foobar'", opt.Str)
	}
	if opt.Nested.Level2_Int != 2 {
		t.Fatalf("Parse struct Int failed, got: '%s' want: 2", opt.Nested.Level2_Int)
	}
}

func TestNestedNilPointerOpt(t *testing.T) {
	opt := &nestedOptionsExample{}
	err := Parse(getter, opt)
	if err == nil {
		t.Fatalf("Parse didn't failed")
	}
	msg := "Nested is nil pointer"
	if err.Error() != msg {
		t.Fatalf("Error message didn't match, got: %v want: %v", err.Error(), msg)
	}
}

type optionsNamedExample struct {
	Str string `name:"str"`
	Int int    `name:"int"`
}

func TestNamedOpt(t *testing.T) {
	opt := &optionsNamedExample{}
	err := Parse(getter, opt)
	if err != nil {
		t.Fatalf("Parse failed with: '%s'", err)
	}
	if opt.Str != "foobar" {
		t.Fatalf("Parse struct Str failed, got: '%s' want: 'foobar'", opt.Str)
	}
	if opt.Int != 1 {
		t.Fatalf("Parse struct Int failed, got: '%s' want: 1", opt.Int)
	}
}

// TODO: Add other types:
// uint, uint32, uint64, int32, int16, float32, float64,
// time.Duration, bool
type optionsTypesExample struct {
	Str    string `name:"str"`
	Int    int    `name:"int"`
	Int64  int64  `name:"int64"`
	UInt16 uint16 `name:"uint16"`
}

func TestTypesOpt(t *testing.T) {
	opt := &optionsTypesExample{}
	err := Parse(getter, opt)
	if err != nil {
		t.Fatalf("Parse failed with: '%s'", err)
	}

	if opt.Str != "foobar" {
		t.Fatalf("Parse struct Str failed, got: '%t' want: 'foobar'", opt.Str)
	}
	if opt.Int != 1 {
		t.Fatalf("Parse struct Int failed, got: '%t' want: 1", opt.Int)
	}
	if opt.Int64 != 1 {
		t.Fatalf("Parse struct Int64 failed, got: '%t' want: 1", opt.Int64)
	}
	if opt.UInt16 != 1 {
		t.Fatalf("Parse struct UInt16 failed, got: '%t' want: 1", opt.UInt16)
	}
}
