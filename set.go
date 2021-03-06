package yec

import "strings"

// SetAppName sets value as new application name
func (y *Yec) SetAppName(value string) {
	y.appName = value
}

// SetAppName sets value as new config file name
func (y *Yec) SetConfigName(value string) {
	y.configName = value
}

func (y *Yec) SetEnvKeyReplacer(r *strings.Replacer) {
	y.envKeyReplacer = r
}
