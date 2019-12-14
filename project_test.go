package libproject

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/BurntSushi/toml"
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

[[build.env]]
name = "JAVA_OPTS"
value = "-Xmx300m"
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
			expected, val.Build.Buildpacks[0].Id)
	}

	expected = "1.0"
	if val.Build.Buildpacks[0].Version != expected {
		t.Fatalf("Expected\n-----\n%#v\n-----\nbut got\n-----\n%#v\n",
			expected, val.Build.Buildpacks[0].Version)
	}

	expected = "https://example.com/buildpack"
	if val.Build.Buildpacks[1].Uri != expected {
		t.Fatalf("Expected\n-----\n%#v\n-----\nbut got\n-----\n%#v\n",
			expected, val.Build.Buildpacks[1].Uri)
	}

	expected = "JAVA_OPTS"
	if val.Build.Env[0].Name != expected {
		t.Fatalf("Expected\n-----\n%#v\n-----\nbut got\n-----\n%#v\n",
			expected, val.Build.Env[0].Name)
	}

	expected = "-Xmx300m"
	if val.Build.Env[0].Value != expected {
		t.Fatalf("Expected\n-----\n%#v\n-----\nbut got\n-----\n%#v\n",
			expected, val.Build.Env[0].Value)
	}
}

func TestFileDoesNotExist(t *testing.T) {
	_, err := ReadProjectDescriptor("/path/that/does/not/exist/project.toml")

	if !os.IsNotExist(err) {
		t.Fatalf("Expected\n-----\n%#v\n-----\nbut got\n-----\n%#v\n",
			"project.toml does not exist error", "no error")
	}
}

func TestReadFile(t *testing.T) {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "project-")
	if err != nil {
		log.Fatal("Cannot create temporary file", err)
	}

	defer os.Remove(tmpFile.Name())

	var text = `
[project]
name = "gallant"
`

	if _, err = tmpFile.Write([]byte(text)); err != nil {
		log.Fatal("Failed to write to temporary file", err)
	}

	// Close the file
	if err := tmpFile.Close(); err != nil {
		log.Fatal(err)
	}

	val, err := ReadProjectDescriptor(tmpFile.Name())

	if err != nil {
		t.Fatal(err)
	}

	var expected = "gallant"
	if val.Project.Name != expected {
		t.Fatalf("Expected\n-----\n%#v\n-----\nbut got\n-----\n%#v\n",
			expected, val.Project.Name)
	}
}

func TestEmtpyEnv(t *testing.T) {
	var testSimple = `
[project]
name = "gallant"
`
	var val ProjectDescriptor
	_, err := toml.Decode(testSimple, &val)
	if err != nil {
		t.Fatal(err)
	}

	expected := 0
	if len(val.Build.Env) != 0 {
		t.Fatalf("Expected\n-----\n%#v\n-----\nbut got\n-----\n%#v\n",
			expected, string(len(val.Build.Env)))
	}

	for _, envVar := range val.Build.Env {
		t.Fatalf("Expected\n-----\n%#v\n-----\nbut got\n-----\n%#v\n",
			"[]", envVar)
	}
}