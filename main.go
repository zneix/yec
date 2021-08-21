package yec

import (
	"bytes"
	"strings"
	//"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

// TODO: read from env and parse it as well
// TODO: make use of appName in multiple config files and env prefix
// TODO: Add logging level and handle fmt.Printf's with debug flags

type Yec struct {
	appName    string
	configName string

	envKeyReplacer StringReplacer

	config map[string]interface{}
	env    map[string][]string
}

// StringReplacer applies a set of replacements to a string.
type StringReplacer interface {
	// Replace returns a copy of s with all replacements performed.
	Replace(s string) string
}

// New returns an initialized Yec instance with default values
// to override these, use methods starting with Set
func New(name string) *Yec {
	return &Yec{
		appName:        name,
		configName:     "config",
		envKeyReplacer: strings.NewReplacer("_", "-"),
		config:         make(map[string]interface{}),
		env:            make(map[string][]string),
	}
}

// ReadConfig reads the configuration file from disk
func (y *Yec) ReadConfig() error {
	// TODO: read config from more paths and merge these into 1 object
	// currently only working directory is read

	fileBytes, err := os.ReadFile(y.configName + ".yaml")
	//fmt.Printf("%# v\n", fileBytes)
	if err != nil {
		return err
	}

	config := make(map[string]interface{})

	err = unmarshalReader(bytes.NewReader(fileBytes), config)
	if err != nil {
		return err
	}

	y.config = config
	//fmt.Printf("%# v\n", config)
	return nil
}

// Unmarshal unmarshals the config into rawVal
// make sure that the tags on the fields of rawVal are properly set
func (y *Yec) Unmarshal(rawVal interface{}) error {
	return decode(y.config, defaultDecoderConfig(rawVal))
}

// unmarshalReader reads data from in and tries to parse it as yaml and unmarshal into c
func unmarshalReader(in io.Reader, c map[string]interface{}) error {
	buf := new(bytes.Buffer)
	buf.ReadFrom(in)

	return yaml.Unmarshal(buf.Bytes(), &c)
}
