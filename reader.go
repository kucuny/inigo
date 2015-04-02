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
	var res []string

	for key, _ := range config.config {
		res = append(res, key)
	}

	return res
}

// Get all keys
func (config *Config) GetAllKeys() map[string][]string {
	var res map[string][]string = make(map[string][]string)

	for section, data := range config.config {
		for key, _ := range data.data {
			res[section] = append(res[section], key)
		}
	}

	return res
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
