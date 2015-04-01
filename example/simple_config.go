package main

import (
	"../../inigo"
	"fmt"
)

func main() {
	config, _ := inigo.LoadConfig("simple_config.ini")
	// fmt.Println(config)
	// fmt.Println(config.GetValue("production", "host"))
	// fmt.Println(config.GetValue("staging", "db.account"))
	// fmt.Println(config.GetValue("staging", "db.kucuny"))
	// fmt.Println(config.GetValue("staging", "db.kucundfdf"))
	// fmt.Println(config.GetValue("production", "host"))

	// fmt.Println(config.GetAllSections())
	fmt.Println(config.GetAllKeys())
}
