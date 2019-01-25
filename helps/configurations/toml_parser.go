package configurations

import (
	"bytes"
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

var _ ConfigurationParser = (*TomlParser)(nil)

type TomlParser struct{}

// Toml data for the specific interface{}(struct)
func (parser *TomlParser) ConfigurationData(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	e := toml.NewEncoder(&buf)
	err := e.Encode(v)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Load toml configuration from file to the specific interface{}(struct)
func (parser *TomlParser) LoadConfiguration(v interface{}, file string) error {
	_, err := toml.DecodeFile(file, v)
	return err
}

// Sava the toml configuration file to the specify file
func (parser *TomlParser) SaveConfiguration(v interface{}, file string) error {
	data, err := parser.ConfigurationData(v)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(file, data, 0644)
}
