package libproject

import (
	"github.com/BurntSushi/toml"
	"testing"
)

func TestDecodeSimple(t *testing.T) {
	var testSimple = `
[project]
name = "gallant"

[build]
exclude = [ "*.jar" ]
`
	var val ProjectDescriptor
	_, err := toml.Decode(testSimple, &val)
	if err != nil {
		t.Fatal(err)
	}

	expected := "gallant"
	if val.Project.Name != "gallant" {
		t.Fatalf("Expected\n-----\n%#v\n-----\nbut got\n-----\n%#v\n",
			expected, val.Project.Name)
	}
}