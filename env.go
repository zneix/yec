package yec

import (
	"fmt"
	"os"
	"strings"
)

func (y *Yec) FindEnv() (envMap map[string]string) {
	envMap = make(map[string]string)

	for _, val := range os.Environ() {
		pair := strings.SplitN(val, "=", 2)
		// os.Environ() should always contain "=", but if it for some reason doesn't return
		// not returning here could cause a panic
		if len(pair) != 2 {
			continue
		}

		varKey := strings.ToLower(pair[0])
		varValue := pair[1]

		ok, cleanKey := y.getKeyFromEnv(varKey)

		if !ok {
			continue
		}

		fmt.Printf("\"%s\": %# v\n", cleanKey, varValue)
		if y.envKeyReplacer != nil {
			cleanKey = y.envKeyReplacer.Replace(cleanKey)
		}

		// Register the variable in our map with env variables
		envMap[cleanKey] = varValue
	}

	fmt.Printf("%# v\n", envMap)

	return
}

// getKeyFromEnv checks if the full environment variable's key starts with y.appName + "_" (prefix)
// if true, returns the key that's after the prefix
func (y *Yec) getKeyFromEnv(varKey string) (bool, string) {
	// Only care about variables that start with the y.appName
	prefix := strings.ToLower(y.appName) + "_"
	key := strings.ToLower(varKey)

	if !strings.HasPrefix(key, prefix) {
		return false, ""
	}

	return true, strings.TrimPrefix(key, prefix)
}

func castMapStringToMapInterface(src map[string]string) map[string]interface{} {
	tgt := map[string]interface{}{}
	for k, v := range src {
		tgt[k] = v
	}
	return tgt
}
