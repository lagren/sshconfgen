package model

type Configuration struct {
	Inventory []Inventory `yaml:"inventory"`
}

type Inventory struct {
	Name string `yaml:"name"`
	Data map[string]string
}

func (i *Inventory) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var raw map[string]string

	err := unmarshal(&raw)
	if err != nil {
		return err
	}

	if name, found := raw["name"]; found {
		i.Name = name
	}

	i.Data = raw

	return nil
}

func (i *Inventory) AsMap() map[string]string {
	return i.Data
}
