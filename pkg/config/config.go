// package that agregates global configuration
package config

import (
	"slices"
	"time"

	"github.com/LCRERGO/GO8EM/pkg/utils/lcg"
)

var config *Config

type Config struct {
	buildtags       []string
	randomGenerator *lcg.LCG
}

// Create a new Config.
func New() *Config {
	return &Config{
		randomGenerator: lcg.New(int(time.Now().UnixNano())),
	}
}

// Get the default instance of the Config.
func Get() *Config {
	if config == nil {
		config = New()
	}

	return config
}

// Add more buildtags to a Config.
func AddBuildTag(config *Config, buildtags ...string) {
	for _, tag := range buildtags {
		if !slices.Contains(config.buildtags, tag) {
			config.buildtags = append(config.buildtags, tag)
		}
	}
}

func HasTag(config *Config, buildtag string) bool {
	return slices.Contains(config.buildtags, buildtag)
}

// Retrives the default randomGenerator
func GetRandomGenerator(config *Config) *lcg.LCG {
	return config.randomGenerator
}
