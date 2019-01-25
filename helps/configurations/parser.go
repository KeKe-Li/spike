package configurations

// ConfigurationParser interface
// Already implemented: TomlParser
type ConfigurationParser interface {
	ConfigurationData(v interface{}) ([]byte, error)
	LoadConfiguration(v interface{}, uri string) error
	SaveConfiguration(v interface{}, uri string) error
}
