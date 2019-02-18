package main

import (
	"os"
	"strings"
)

type tmplValues struct {
	Env map[string]string
}

func loadTmplValues(envPrefix string) tmplValues {
	if envPrefix != "" {
		envPrefix = envPrefix + "_"
	}
	envVars := make(map[string]string)
	for _, env := range os.Environ() {
		splitEnv := strings.SplitN(env, "=", 2)
		if strings.HasPrefix(splitEnv[0], envPrefix) {
			envVars[strings.TrimPrefix(splitEnv[0], envPrefix)] = splitEnv[1]
		}
	}
	return tmplValues{envVars}
}
