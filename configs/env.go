package configs

import (
	"fmt"
	"os"
	"sync"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type envConfig struct {
	PORT int `env:"PORT" envDefault:"8080"`
}

var config *envConfig
var lock sync.Mutex

func LoadEnv() *envConfig {
	lock.Lock()
	defer lock.Unlock()

	s, _ := os.Getwd()
	fmt.Println(s)

	if config != nil {
		return config
	}

	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
	}

	config = &envConfig{}
	if err := env.Parse(config); err != nil {
		fmt.Printf("%+v\n", err)
	}
	return config

}
