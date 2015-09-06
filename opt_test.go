package cfg

import (
	"os"
	"testing"
)

func TestStringOpt(t *testing.T) {
	opt := Option{Name: "test"}
	e := "Option 'test' (default: <nil>, required: false"
	if s := opt.String(); s != e {
		t.Fatalf("String representation didn't match, got: '%s', want: '%s'", s, e)
	}
}

func TestEmptyOpt(t *testing.T) {
	opt := Option{Name: "test"}
	if v, _ := opt.Get(emptyGetter); v != nil {
		t.Fatalf("Empty option didn't return nil, got: '%s'", v)
	}
}

func TestDefaultValueOpt(t *testing.T) {
	opt := Option{Name: "test", DefaultValue: "foobar"}
	if v, _ := opt.Get(emptyGetter); v != "foobar" {
		t.Fatalf("Couldn't get default value, got: '%s'", v)
	}
}

func TestEnvValueOpt(t *testing.T) {
	os.Setenv("CFG_TEST", "foobar")
	opt := Option{Name: "test", EnvVar: "CFG_TEST"}
	if v, _ := opt.Get(emptyGetter); v != "foobar" {
		t.Fatalf("Couldn't get environment value, got: '%s'", v)
	}
}

func TestRequiredValueOpt(t *testing.T) {
	opt := Option{Name: "test", Required: true}
	if _, err := opt.Get(emptyGetter); err == nil {
		t.Fatal("Empty required option didn't return an error")
	}
}
