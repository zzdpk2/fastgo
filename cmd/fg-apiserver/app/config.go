package app

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

const (
	defaultHomeDir = ".fastgo"

	defaultConfigName = "fg-apiserver.yaml"
)

func onInitialize() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		for _, dir := range searchDirs() {
			viper.AddConfigPath(dir)
		}

		viper.SetConfigType("yaml")

		viper.SetConfigName(defaultConfigName)
	}

	setupEnvironmentVariables()
	_ = viper.ReadInConfig()
}

func searchDirs() []string {
	homeDir, err := os.UserHomeDir()
	cobra.CheckErr(err)
	return []string{filepath.Join(homeDir, defaultHomeDir), "."}
}

func setupEnvironmentVariables() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("FASTGO")
	replacer := strings.NewReplacer(".", "_", "-", "_")
	viper.SetEnvKeyReplacer(replacer)
}

func filePath() string {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)
	return filepath.Join(home, defaultHomeDir, defaultConfigName)
}
