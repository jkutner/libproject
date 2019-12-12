package libproject

type Build struct {
	Include []string `toml:"include"`
	Exclude []string `toml:"exclude"`
}

type Project struct {
	Name string `toml:"name"`
}

type ProjectDescriptor struct {
	Project  Project                `toml:"project"`
	Build    Build                  `toml:"build"`
	Metadata map[string]interface{} `toml:"metadata"`
}
