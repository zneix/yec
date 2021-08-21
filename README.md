# Yec: YAML-env-config wrapper for Golang

Yec is a simple library that reads data from yaml files and environment variables unmarshaling it into custom structs.
It aims to be as lightweight as possible, relying on minimal amount of dependencies, while providing unified functionality between configuration files and environment variables.

Simple example:

```go
package main

import (
	"fmt"

	"github.com/zneix/yec"
)

type myConfig struct {
	BaseURL          string `mapstructure:"base-url"`
	MaxContentLength uint64 `mapstructure:"max-content-length"`
}

func main() {
	y := yec.New("myapplication")
	y.ReadInConfig()

	var cfg myConfig
	y.Unmarshal(&cfg)

	fmt.Printf("%# v\n", cfg)
}
```

Inspired by [spf13/viper](https://github.com/spf13/viper).
