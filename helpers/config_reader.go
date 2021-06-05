package helpers

import "github.com/spf13/viper"

// ConfigReader interface helps me to abstract away config retrieval to better unit test the tool
type ConfigReader interface {
	GetString(key string) string
	GetStringSlice(key string) []string
	GetInt(key string) int
	IsSet(key string) bool
}

// ViperConfigReader is a concrete implementation of the ConfigReader interface
type ViperConfigReader struct {
}

// GetString returns a string type config entry
func (v ViperConfigReader) GetString(key string) string {
	return viper.GetString(key)
}

// GetStringSlice returns a string slice type config entry
func (v ViperConfigReader) GetStringSlice(key string) []string {
	return viper.GetStringSlice(key)
}

// GetInt returns an int type config entry
func (v ViperConfigReader) GetInt(key string) int {
	return viper.GetInt(key)
}

// IsSet confirms the existence of a config key
func (v ViperConfigReader) IsSet(key string) bool {
	return viper.IsSet(key)
}

//go:generate mockgen -source=config_reader.go -package=helpers -destination=mock_config_reader.go
