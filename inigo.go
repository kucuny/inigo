package inigo

import (
	"sync"
)

const (
	LINE_BREAK        = "\n"
	SECTION_SEPERATOR = ":"
	KEY_SEPERATOR     = "."
)

const (
	REGEX_SECTION             = `\[(.*)\]`
	REGEX_KEY_VALUE           = `(\S*)=(.*)`
	REGEX_INHERITANCE_SECTION = `(.*)(?:\:(.*))`
)

const DEFAULT_SECTION = "default"

type ConfigData struct {
	inheritSection string
	data           map[string]string
}

type Config struct {
	mu       sync.RWMutex
	filename string
	config   map[string]ConfigData
	sections []string
	keys     map[string][]string
}
