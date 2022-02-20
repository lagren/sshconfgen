package model

type Template struct {
	Name     string              `yaml:"name"`
	Filename string              `yaml:"filename"`
	Config   map[string]string   `yaml:"config"`
	Filters  []map[string]string `yaml:"filters"`
}
