package discover

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

func FindUprojectInPath(projectPath string) (string, error) {
	return FindFilePath(projectPath, "*.uproject", false)
}

func FindUe5EnginePath() (string, error) {
	knownPaths := [...]string{
		"/Users/Shared/Epic Games/",
		"/Users/Shared/",
	}

	filenamePatterns := [...]string{
		"Engine",
	}

	for _, path := range knownPaths {
		for _, pattern := range filenamePatterns {
			matchedPath, err := FindFolderPath(path, pattern, 3)

			if err == nil {
				return matchedPath, nil
			}
		}
	}

	return "", fmt.Errorf("Unreal Engine could not be found!")
}

func FindUnrealEditorPath(ue5Root string) (string, error) {
	searchPath := filepath.Join(ue5Root, "Binaries/Mac/UnrealEditor.app/")

	return FindFilePath(searchPath, "UnrealEditor", true)
}

func FindBuildScriptPath(ue5Root string) (string, error) {
	searchPath := filepath.Join(ue5Root, "Build/BatchFiles/Mac")

	return FindFilePath(searchPath, "Build.sh", true)
}

func FindBuildToolPath(ue5Root string) (string, error) {
	searchPath := filepath.Join(ue5Root, "Binaries/DotNET/UnrealBuildTool")

	return FindFilePath(searchPath, "UnrealBuildTool", true)
}

func FindFilePath(rootPath string, pattern string, isRecursive bool) (string, error) {
	matchedPath := ""

	err := filepath.WalkDir(rootPath, func(path string, d fs.DirEntry, errWalk error) error {
		if d == nil || path == rootPath {
			return nil
		}

		if d.IsDir() && !isRecursive {
			return filepath.SkipDir
		}

		isMatch, _ := filepath.Match(pattern, d.Name())

		if isMatch {
			matchedPath = path
			return filepath.SkipAll
		}

		return errWalk
	})

	if err != nil {
		return matchedPath, err
	}

	if matchedPath == "" {
		err = fmt.Errorf("No match found for %s in %s", pattern, rootPath)
	}

	return matchedPath, err
}

func FindFolderPath(rootPath string, pattern string, maxSearchDepth int) (string, error) {
	var matchedPath string

	pathSeparatorStr := string(filepath.Separator)
	rootPathDepth := strings.Count(rootPath, pathSeparatorStr)

	err := filepath.WalkDir(rootPath, func(path string, d fs.DirEntry, errWalk error) error {
		if errWalk != nil {
			return errWalk
		}

		if d == nil {
			return nil
		}

		pathDepth := strings.Count(path, pathSeparatorStr)

		if (pathDepth - rootPathDepth) > maxSearchDepth {
			return nil
		}

		if d.IsDir() {
			matched, _ := filepath.Match(pattern, d.Name())

			if matched {
				matchedPath = path
				return fs.SkipAll
			}
		}

		return nil
	})

	if err != nil {
		return matchedPath, err
	}

	if matchedPath == "" {
		err = fmt.Errorf("No match found for %s in %s", pattern, rootPath)
	}

	return matchedPath, err
}
