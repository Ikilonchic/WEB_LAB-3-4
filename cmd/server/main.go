package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"

	"github.com/ikilonchic/WEB_LAB-3-4/internal/server"
)

var (
	configFile string
)

func init() {
	flag.StringVar(&configFile, "-config-path", "./configs/server.toml", "Path to config file (*.toml)")
}

func main() {
	flag.Parse()

	config := server.NewConfig()
	if _, err := toml.DecodeFile(configFile, config); err != nil {
		log.Fatal(err)
	}

	serv := server.New(config)

	if err := serv.Start(); err != nil {
		log.Fatal(err)
	}
}