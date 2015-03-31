package inigo

import (
	"bufio"
	"os"
	"regexp"
)

const (
	PARSE_PROCESS_INIT      = 0
	PARSE_PROCESS_SECTION   = 1
	PARSE_PROCESS_KEY_VALUE = 2
)

// Load INI config
func LoadConfig(filename string) (config *Config, err error) {
	config = new(Config)
	config.filename = filename
	config.config = make(map[string]ConfigData)
	config.keys = make(map[string][]string)

	c := ConfigData{
		inheritSection: "",
	}

	config.config[DEFAULT_SECTION] = c

	config.parseIni()

	return config, nil
}

// Reload config file
func (config *Config) ReloadConfig() (err error) {
	c, err := LoadConfig(config.filename)
	c.parseIni()

	config = c

	return err
}

// Create config data
func createConfigInitData(inherit string) ConfigData {
	resInherit := DEFAULT_SECTION

	if inherit != "" {
		resInherit = inherit
	}

	configData := ConfigData{
		inheritSection: resInherit,
		data:           make(map[string]string),
	}

	return configData
}

// Get section and inherit section
func getSection(source string) (section, inheritSection string) {
	sectionReg, _ := regexp.Compile(REGEX_SECTION)

	if ok := sectionReg.MatchString(source); ok {
		section = sectionReg.FindStringSubmatch(source)[1]

		inheritSectionReg, _ := regexp.Compile(REGEX_INHERITANCE_SECTION)

		inheritSection = DEFAULT_SECTION

		if ok := inheritSectionReg.MatchString(section); ok {
			inheritSection = inheritSectionReg.FindStringSubmatch(section)[2]
			section = inheritSectionReg.FindStringSubmatch(section)[1]
		}
	}

	return
}

// Get key-value in section
func getKeyValue(source string) (key, value string) {
	reg, _ := regexp.Compile(REGEX_KEY_VALUE)

	if ok := reg.MatchString(source); ok {
		key = reg.FindStringSubmatch(source)[1]
		value = reg.FindStringSubmatch(source)[2]
	}

	return
}

// Parse INI file and save values
func (config *Config) parseIni() (err error) {
	f, err := os.Open(config.filename)
	defer f.Close()

	config.mu.Lock()
	defer config.mu.Unlock()

	reader := bufio.NewReader(f)
	scanner := bufio.NewScanner(reader)

	var processing = PARSE_PROCESS_SECTION
	var currentSection string

	for scanner.Scan() {
		if scanner.Err() != nil {
			break
		}

		line := scanner.Text()

		// Section
		section, inheritSection := getSection(line)
		if len(section) > 0 && len(inheritSection) > 0 {
			processing = PARSE_PROCESS_KEY_VALUE
			currentSection = section
			c := createConfigInitData(inheritSection)
			config.config[currentSection] = c
		}

		// Key-Value
		if processing == PARSE_PROCESS_KEY_VALUE {
			key, value := getKeyValue(line)
			if len(key) > 0 {
				config.config[currentSection].data[key] = value
			}
		}
	}

	return err
}
