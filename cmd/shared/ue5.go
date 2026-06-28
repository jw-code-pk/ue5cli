package shared

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func GetUe5ProjectName(projectFilePath string) (string, error) {
	info, err := os.Stat(projectFilePath)

	if err != nil {
		return "", err
	}

	projectName := info.Name()
	projectName = strings.TrimSuffix(projectName, ".uproject")

	return projectName, nil
}

func RunUe5Cmd(cmd *exec.Cmd) error {

	stdoutPipe, err := cmd.StdoutPipe()

	if err != nil {
		return err
	}

	err = cmd.Start()

	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(stdoutPipe)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("[UE5LOG] %s\n", line)
	}

	err = scanner.Err()

	if err != nil {
		return err
	}

	err = cmd.Wait()

	if err != nil {
		return err
	}

	return nil
}
