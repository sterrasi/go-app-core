package core

type Configuration struct {
	profiles []string
}

func NewConfiguration() *Configuration {
	return &Configuration{}
}

type FileConfig interface {
	Value() (*string, error)
	IntValue() (*int, error)
	FloatValue() (*float32, error)
	BoolValue() (*bool, error)
}

// CompositeFileConfig a file configuration that delegates to a list of file configurations, and returns the first match
type CompositeFileConfig struct {
}
