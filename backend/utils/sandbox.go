package utils

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

// ExecutionResult represents the result of code execution
type ExecutionResult struct {
	Output  string `json:"output"`
	Error   string `json:"error"`
	Success bool   `json:"success"`
	Time    int    `json:"time"` // execution time in milliseconds
}

// ExecuteCode safely executes Go code in a sandboxed environment
func ExecuteCode(code string, timeout time.Duration) (*ExecutionResult, error) {
	startTime := time.Now()

	// Create a temporary directory for this execution
	tempDir := filepath.Join(os.TempDir(), "gohacker-"+uuid.New().String())
	if err := os.MkdirAll(tempDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create temp directory: %w", err)
	}
	defer os.RemoveAll(tempDir)

	// Write the code to a temporary file
	codeFile := filepath.Join(tempDir, "main.go")
	if err := os.WriteFile(codeFile, []byte(code), 0644); err != nil {
		return nil, fmt.Errorf("failed to write code file: %w", err)
	}

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// First, try to build the code
	buildCmd := exec.CommandContext(ctx, "go", "build", "-o", filepath.Join(tempDir, "program"), codeFile)
	buildCmd.Dir = tempDir
	var buildOut, buildErr bytes.Buffer
	buildCmd.Stdout = &buildOut
	buildCmd.Stderr = &buildErr

	if err := buildCmd.Run(); err != nil {
		return &ExecutionResult{
			Output:  buildOut.String(),
			Error:   buildErr.String(),
			Success: false,
			Time:    int(time.Since(startTime).Milliseconds()),
		}, nil
	}

	// If build succeeded, run the program
	runCmd := exec.CommandContext(ctx, filepath.Join(tempDir, "program"))
	runCmd.Dir = tempDir
	var runOut, runErr bytes.Buffer
	runCmd.Stdout = &runOut
	runCmd.Stderr = &runErr

	err := runCmd.Run()

	result := &ExecutionResult{
		Output:  runOut.String(),
		Error:   runErr.String(),
		Success: err == nil,
		Time:    int(time.Since(startTime).Milliseconds()),
	}

	// Check if context was cancelled (timeout)
	if ctx.Err() == context.DeadlineExceeded {
		result.Error = "Execution timeout exceeded"
		result.Success = false
	}

	return result, nil
}

// ValidateGoCode performs basic validation on Go code
func ValidateGoCode(code string) error {
	if len(code) == 0 {
		return fmt.Errorf("code cannot be empty")
	}

	if len(code) > 10000 {
		return fmt.Errorf("code exceeds maximum length of 10000 characters")
	}

	// Check for package declaration
	if !bytes.Contains([]byte(code), []byte("package main")) {
		return fmt.Errorf("code must contain 'package main'")
	}

	return nil
}
