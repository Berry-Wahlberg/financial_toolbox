package python

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"time"
)

type PythonManager struct {
	timeout time.Duration
}

func NewPythonManager() *PythonManager {
	return &PythonManager{
		timeout: 10 * time.Second,
	}
}

func (pm *PythonManager) ExecuteScript(scriptPath string, inputData string) (string, error) {
	cmd := exec.Command("python", scriptPath)
	cmd.Stdin = bytes.NewBufferString(inputData)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Start(); err != nil {
		return "", fmt.Errorf("failed to start Python process: %w", err)
	}

	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()

	select {
	case err := <-done:
		if err != nil {
			errMsg := stderr.String()
			if errMsg != "" {
				return "", fmt.Errorf("Python script error: %s", errMsg)
			}
			return "", fmt.Errorf("Python process failed: %w", err)
		}
		return stdout.String(), nil
	case <-time.After(pm.timeout):
		if cmd.Process != nil {
			cmd.Process.Kill()
		}
		return "", fmt.Errorf("Python script execution timeout")
	}
}

func (pm *PythonManager) ExecuteWithJSON(scriptPath string, data interface{}) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("failed to marshal input data: %w", err)
	}

	return pm.ExecuteScript(scriptPath, string(jsonData))
}

func (pm *PythonManager) SetTimeout(timeout time.Duration) {
	pm.timeout = timeout
}
