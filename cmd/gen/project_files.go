package gen

import (
	"fmt"
	"os/exec"
	"ue5cli/cmd/config"
	"ue5cli/cmd/shared"
)

func GenerateProjectFiles() error {
	fmt.Println("📋 Generating Unreal project files...")

	buildToolPath := config.Ue5BuildToolPath
	projectFilePath := config.Ue5ProjectFilePath
	targetName, err := shared.GetUe5ProjectName(projectFilePath)

	if err != nil {
		return err
	}

	targetName += "Editor"

	// TODO: there are other flags to configure here - check the docs! (also are all these ones needed?)
	buildToolCmd := exec.Command(
		buildToolPath,
		"-ProjectFiles",
		"-Project="+projectFilePath,
		"-Game",
		"-Engine",
		targetName,
		"Mac",
		"Development",
	)

	shared.RunUe5Cmd(buildToolCmd)

	return nil
}
