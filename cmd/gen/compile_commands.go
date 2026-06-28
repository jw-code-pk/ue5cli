package gen

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"ue5cli/cmd/config"
	"ue5cli/cmd/shared"
)

func GenerateCompileCommands() error {
	fmt.Println("📋 Generating compile_commands...")
	projectFilePath := config.Ue5ProjectFilePath
	targetName, err := shared.GetUe5ProjectName(projectFilePath)

	if err != nil {
		return err
	}

	buildToolCmd := exec.Command(
		config.Ue5BuildToolPath,
		"-mode=GenerateClangDatabase",
		"-project="+projectFilePath,
		"-game",
		"-engine",
		targetName+"Editor",
		"Mac",
		"Development",
	)

	shared.RunUe5Cmd(buildToolCmd)

	outputFileName := "compile_commands.json"
	outputFilePath := filepath.Join(config.Ue5ProjectDir, outputFileName)
	_, err = os.Lstat(outputFilePath)
	if err == nil {
		fmt.Printf("'%s' exists. Removing...\n", outputFilePath)
		err = os.Remove(outputFilePath)
		if err != nil {
			return err
		}
	} else if !os.IsNotExist(err) {
		return err
	}

	inputFilePath := filepath.Join(config.Ue5RootPath, outputFileName)

	inputFile, err := os.Open(inputFilePath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	_, err = io.Copy(outputFile, inputFile)
	if err != nil {
		return err
	}

	fmt.Printf("🛠️ Created '%s'\n", outputFilePath)

	return nil
}
