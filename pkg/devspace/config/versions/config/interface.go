package config

// Config is the interface for each config version
type Config interface {
	GetVersion() string
	Upgrade() (Config, error)
}

// New creates a new config
type New func() Config

// Variables strips all information from the config except variables
type Variables func(data map[interface{}]interface{}) (map[interface{}]interface{}, error)

// Commands strips all information from the config except commands
type Commands func(data map[interface{}]interface{}) (map[interface{}]interface{}, error)

// Profile loads a certain profile with the base config
type Profile func(data map[interface{}]interface{}, profile string) (map[interface{}]interface{}, error)
