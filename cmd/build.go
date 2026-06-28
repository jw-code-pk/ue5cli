/*
Copyright © 2026 John-Evans Wagenaar <jewagenaar@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"os/exec"
	"ue5cli/cmd/config"
	"ue5cli/cmd/shared"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Run the Unreal build script on the project.",
	RunE: func(cmd *cobra.Command, args []string) error {
		buildScriptPath := config.Ue5BuildScriptPath
		projectFilePath := config.Ue5ProjectFilePath

		fmt.Printf("🛠️ Run build: %s %s\n", buildScriptPath, projectFilePath)

		projectName, err := shared.GetUe5ProjectName(projectFilePath)

		if err != nil {
			return err
		}

		// TODO: make all this configurable via flags (Target, Platform, BuildType)
		targetName := projectName + "Editor"
		buildCmd := exec.Command(buildScriptPath, targetName, "Mac", "Development", projectFilePath)

		err = shared.RunUe5Cmd(buildCmd)

		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
