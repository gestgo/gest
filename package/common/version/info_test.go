package version

import (
	"strings"
	"testing"
)

type testMeta struct {
	Env     string `json:"env"`
	Cluster string `json:"cluster"`
}

func TestNew_NoExtra(t *testing.T) {
	info := New[any](nil)
	if info.AppName != AppName {
		t.Errorf("expected AppName=%s, got %s", AppName, info.AppName)
	}
	if info.GoVersion == "" {
		t.Error("GoVersion should be populated")
	}
}

func TestNew_StructExtra(t *testing.T) {
	info := New(testMeta{Env: "prod", Cluster: "blue"})
	if info.Extra.Env != "prod" {
		t.Errorf("expected env=prod, got %s", info.Extra.Env)
	}
	if info.Extra.Cluster != "blue" {
		t.Errorf("expected cluster=blue, got %s", info.Extra.Cluster)
	}
}

func TestNew_MapExtra(t *testing.T) {
	info := New(map[string]any{"service": "billing", "tenant": "acme"})
	if info.Extra["service"] != "billing" {
		t.Errorf("expected service=billing, got %v", info.Extra["service"])
	}
}

func TestString_ContainsFields(t *testing.T) {
	orig := Version
	Version = "1.2.3"
	t.Cleanup(func() { Version = orig })

	box := New(testMeta{Env: "prod", Cluster: "blue"}).String()
	for _, want := range []string{"1.2.3", "prod", "blue", "extra"} {
		if !strings.Contains(box, want) {
			t.Errorf("box missing %q", want)
		}
	}
}

func TestExtraLines_Primitive(t *testing.T) {
	lines := extraLines("hello")
	if len(lines) != 1 || lines[0] != "hello" {
		t.Errorf("unexpected lines: %v", lines)
	}
}

func TestExtraLines_Nil(t *testing.T) {
	if extraLines(nil) != nil {
		t.Error("nil extra should return nil lines")
	}
}
