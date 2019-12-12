package libproject

type Buildpack struct {
	Id      string `toml:"id"`
	Version string `toml:"version"`
	Uri     string `toml:"uri"`
}

type Build struct {
	Include []string `toml:"include"`
	Exclude []string `toml:"exclude"`
	Buildpacks []Buildpack `toml:"buildpacks"`
}

type Project struct {
	Name string `toml:"name"`
}

type ProjectDescriptor struct {
	Project  Project                `toml:"project"`
	Build    Build                  `toml:"build"`
	Metadata map[string]interface{} `toml:"metadata"`
}
