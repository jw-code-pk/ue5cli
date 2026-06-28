package config

import (
	"fmt"
	"os"
	"path/filepath"
	"ue5cli/cmd/discover"

	"github.com/spf13/viper"
)

var (
	configFileName = ".ue5-cli"
	configFileExt  = "yaml"
)

func InitConfig() {
	home, _ := os.UserHomeDir()

	// setup default config
	viper.AddConfigPath(home)
	viper.SetConfigName(configFileName)
	viper.SetConfigType(configFileExt)
	viper.SetDefault("projectDir", ".")

	viper.ReadInConfig()

	RefreshGlobalConfig()
}

func WriteConfigFile() {
	home, _ := os.UserHomeDir()
	configFilePath := filepath.Join(home, configFileName+"."+configFileExt)
	_, err := os.Stat(configFilePath)

	if err != nil && os.IsNotExist(err) {
		// create the config file if it wasn't there
		configFile, err := os.OpenFile(configFilePath, os.O_RDONLY|os.O_CREATE, 0644)
		if err != nil {
			fmt.Printf("[ERROR] Could not write config file! %s\n", err)
		}
		configFile.Close()
	}

	err = viper.WriteConfig()

	if err != nil {
		fmt.Printf("[ERROR] Could not write config file! %s\n", err)
	}

	RefreshGlobalConfig()
}

func RefreshGlobalConfig() {
	projectDir := viper.GetString("projectDir")

	if projectDir == "." {
		workingDir, err := os.Getwd()
		if err != nil {
			fmt.Printf("[ERROR] Could not determine working directory!\n%s\n", err)
		}
		projectDir = workingDir
	}

	fmt.Printf("Using project folder '%s'\n", projectDir)

	projectPath, err := discover.FindUprojectInPath(projectDir)

	if err != nil {
		fmt.Printf("[ERROR] Could not find an Unreal project in %s! %s\n", projectDir, err)
	}

	Ue5ProjectDir = projectDir
	Ue5ProjectFilePath = projectPath
	Ue5BuildScriptPath = viper.GetString("buildScriptPath")
	Ue5BuildToolPath = viper.GetString("buildToolPath")
	Ue5EditorPath = viper.GetString("editorPath")
	Ue5EnginePath = viper.GetString("enginePath")
	Ue5RootPath = viper.GetString("rootPath")
}
