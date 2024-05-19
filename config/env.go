package config

import (
	"fmt"
	"os"
)

func GetEnv(name string) string {
	env, found := os.LookupEnv(name)

	if !found {
		panic(fmt.Sprintf("Environment %s not found", name))
	}

	return env
}
