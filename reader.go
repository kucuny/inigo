package inigo

import (
	"errors"
)

// Get config file name
func (config *Config) GetConfigFilename() string {
	return config.filename
}

// Get all config sections
func (config *Config) GetAllSections() []string {
	return nil
}

// Get all keys
func (config *Config) GetAllKeys() map[string][]string {
	return nil
}

// Get value
func (config *Config) GetValue(section, key string) (string, error) {
	var processing bool = true
	var currentSection string = section
	var res string

	for processing && len(currentSection) > 0 {
		res = config.config[currentSection].data[key]

		if len(res) > 0 {
			processing = false
			break
		} else {
			currentSection = config.config[currentSection].inheritSection
		}
	}

	if len(res) == 0 {
		return res, errors.New("Has no key")
	}

	return res, nil
}
