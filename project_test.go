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

[[build.buildpacks]]
id = "example/lua"
version = "1.0"

[[build.buildpacks]]
uri = "https://example.com/buildpack"
`
	var val ProjectDescriptor
	_, err := toml.Decode(testSimple, &val)
	if err != nil {
		t.Fatal(err)
	}

	var expected string

	expected = "gallant"
	if val.Project.Name != expected {
		t.Fatalf("Expected\n-----\n%#v\n-----\nbut got\n-----\n%#v\n",
			expected, val.Project.Name)
	}

	expected = "example/lua"
	if val.Build.Buildpacks[0].Id != expected {
		t.Fatalf("Expected\n-----\n%#v\n-----\nbut got\n-----\n%#v\n",
			expected, val.Project.Name)
	}

	expected = "1.0"
	if val.Build.Buildpacks[0].Version != expected {
		t.Fatalf("Expected\n-----\n%#v\n-----\nbut got\n-----\n%#v\n",
			expected, val.Project.Name)
	}

	expected = "https://example.com/buildpack"
	if val.Build.Buildpacks[1].Uri != expected {
		t.Fatalf("Expected\n-----\n%#v\n-----\nbut got\n-----\n%#v\n",
			expected, val.Project.Name)
	}
}
