/*
Copyright © 2026 John-Evans Wagenaar <jewagenaar@gmail.com>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"ue5cli/cmd/gen"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate various project resources for an UE5 project",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if cmd.Flags().NFlag() == 0 {
			return fmt.Errorf("No resource specified for the 'generate' command.")
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		hasCommands, err := cmd.Flags().GetBool("commands")
		if err != nil {
			return err
		}
		if hasCommands {
			genErr := gen.GenerateCompileCommands()
			if genErr != nil {
				return genErr
			}
		}

		hasProject, err := cmd.Flags().GetBool("project")
		if err != nil {
			return err
		}
		if hasProject {
			genErr := gen.GenerateProjectFiles()
			if genErr != nil {
				return genErr
			}
		}

		if cmd.Flag("cpp").Changed {

			classType := gen.BaseClass

			if hasActor, _ := cmd.Flags().GetBool("actor"); hasActor {
				classType = gen.ActorClass
			} else if hasComponent, _ := cmd.Flags().GetBool("component"); hasComponent {
				classType = gen.ComponentClass
			} else if hasObject, _ := cmd.Flags().GetBool("object"); hasObject {
				classType = gen.ObjectClass
			} else if hasData, _ := cmd.Flags().GetBool("data"); hasData {
				classType = gen.DataClass
			}

			cppClassName, err := cmd.Flags().GetString("cpp")
			if err != nil {
				return err
			}
			folderName, err := cmd.Flags().GetString("folder")
			genErr := gen.GenerateCppClass(cppClassName, folderName, classType)
			if genErr != nil {
				return genErr
			}

			gen.GenerateCompileCommands()
		}
		return nil
	},
}

func init() {
	generateCmd.Flags().String("cpp", "MyClass", "Generate a C++ class in the project's source folder.")
	generateCmd.Flags().Bool("actor", false, "Output an Unreal actor class.")
	generateCmd.Flags().Bool("component", false, "Output an Unreal component class.")
	generateCmd.Flags().Bool("object", false, "Output an Unreal object class.")
	generateCmd.Flags().Bool("data", false, "Output an Unreal data asset class.")

	generateCmd.Flags().String("folder", "", "Relative output folder for project resources.")
	generateCmd.Flags().Bool("commands", false, "Generate a compile_commands file for clangd.")
	generateCmd.Flags().Bool("project", false, "Generate the project files.")
	rootCmd.AddCommand(generateCmd)
}
