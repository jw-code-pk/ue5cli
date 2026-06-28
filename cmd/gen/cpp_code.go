package gen

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"ue5cli/cmd/config"
	"ue5cli/cmd/gen/cpp_templates"
	"ue5cli/cmd/shared"
)

type ClassType int

const (
	BaseClass ClassType = iota
	ActorClass
	ComponentClass
	ObjectClass
	DataClass
)

type CppClassInfo struct {
	Name        string
	FolderPath  string
	ApiName     string
	ProjectName string
}

func GenerateCppClass(cppClassName string, folderPath string, classType ClassType) error {
	projectName, err := shared.GetUe5ProjectName(config.Ue5ProjectFilePath)

	if err != nil {
		return err
	}

	formattedFolderPath := folderPath

	if formattedFolderPath != "" {
		formattedFolderPath += "/"
	}

	apiName := strings.ToUpper(projectName) + "_API"

	classInfo := CppClassInfo{
		Name:        cppClassName,
		FolderPath:  formattedFolderPath,
		ApiName:     apiName,
		ProjectName: projectName,
	}

	htd := cpp_templates.BaseHTD
	std := cpp_templates.BaseSTD

	switch classType {
	case ActorClass:
		{
			htd = cpp_templates.ActorHTD
			std = cpp_templates.ActorSTD
		}
	case ComponentClass:
		{
			htd = cpp_templates.ComponentHTD
			std = cpp_templates.ComponentSTD
		}
	case ObjectClass:
		{
			htd = cpp_templates.ObjectHTD
			std = cpp_templates.ObjectSTD
		}
	case DataClass:
		{
			htd = cpp_templates.DataHTD
			std = cpp_templates.DataSTD
		}
	}

	err = generateHeaderFile(htd, classInfo)

	if err != nil {
		return err
	}

	err = generateSourceFile(std, classInfo)

	if err != nil {
		return err
	}

	return nil
}

func generateHeaderFile(headerTemplateData string, classInfo CppClassInfo) error {
	headerTemplate, err := template.New("CppHeader").Parse(headerTemplateData)

	if err != nil {
		return err
	}

	projectDirectory := config.Ue5ProjectDir
	headerPath := filepath.Join(
		projectDirectory,
		"Source",
		classInfo.ProjectName,
		"Public",
		classInfo.FolderPath)

	err = os.MkdirAll(headerPath, 0750)

	if err != nil {
		return err
	}

	headerFilePath := filepath.Join(headerPath, classInfo.Name+".h")

	_, err = os.Lstat(headerFilePath)

	if err == nil {
		return fmt.Errorf("The header file '%s' already exists.\n", headerFilePath)
	}

	fmt.Printf("Creating header file in '{%s}'\n", headerFilePath)

	headerFile, err := os.Create(headerFilePath)
	if err != nil {
		return err
	}
	defer headerFile.Close()

	err = headerTemplate.Execute(headerFile, classInfo)

	if err != nil {
		return err
	}
	return nil
}

func generateSourceFile(sourceTemplateData string, classInfo CppClassInfo) error {
	sourceTemplate, err := template.New("CppSource").Parse(sourceTemplateData)

	if err != nil {
		return err
	}

	projectDirectory := config.Ue5ProjectDir
	sourcePath := filepath.Join(
		projectDirectory,
		"Source",
		classInfo.ProjectName,
		"Private",
		classInfo.FolderPath)

	err = os.MkdirAll(sourcePath, 0750)

	if err != nil {
		return err
	}

	sourceFilePath := filepath.Join(sourcePath, classInfo.Name+".cpp")

	_, err = os.Lstat(sourceFilePath)

	if err == nil {
		return fmt.Errorf("The source file '%s' already exists.\n", sourceFilePath)
	}

	fmt.Printf("Creating source file in '{%s}'\n", sourceFilePath)

	sourceFile, err := os.Create(sourceFilePath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	err = sourceTemplate.Execute(sourceFile, classInfo)

	if err != nil {
		return err
	}
	return nil
}
