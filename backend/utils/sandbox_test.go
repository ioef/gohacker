package utils

import (
	"strings"
	"testing"
	"time"
)

func TestValidateGoCode(t *testing.T) {
	tests := []struct {
		name    string
		code    string
		wantErr bool
	}{
		{
			name: "valid code",
			code: `package main
import "fmt"
func main() {
	fmt.Println("Hello")
}`,
			wantErr: false,
		},
		{
			name:    "empty code",
			code:    "",
			wantErr: true,
		},
		{
			name:    "missing package main",
			code:    `import "fmt"`,
			wantErr: true,
		},
		{
			name:    "code too long",
			code:    strings.Repeat("a", 10001),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateGoCode(tt.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateGoCode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestExecuteCode(t *testing.T) {
	tests := []struct {
		name           string
		code           string
		expectSuccess  bool
		expectedOutput string
	}{
		{
			name: "simple hello world",
			code: `package main
import "fmt"
func main() {
	fmt.Println("Hello, World!")
}`,
			expectSuccess:  true,
			expectedOutput: "Hello, World!",
		},
		{
			name: "syntax error",
			code: `package main
func main() {
	fmt.Println("Missing import")
}`,
			expectSuccess: false,
		},
		{
			name: "runtime error",
			code: `package main
func main() {
	panic("test panic")
}`,
			expectSuccess: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ExecuteCode(tt.code, 5*time.Second)
			if err != nil {
				t.Fatalf("ExecuteCode() error = %v", err)
			}

			if result.Success != tt.expectSuccess {
				t.Errorf("Expected success=%v, got %v", tt.expectSuccess, result.Success)
			}

			if tt.expectSuccess && !strings.Contains(result.Output, tt.expectedOutput) {
				t.Errorf("Expected output to contain %q, got %q", tt.expectedOutput, result.Output)
			}
		})
	}
}

func TestExecuteCodeTimeout(t *testing.T) {
	code := `package main
import "time"
func main() {
	time.Sleep(10 * time.Second)
}
`

	result, err := ExecuteCode(code, 1*time.Second)
	if err != nil {
		t.Fatalf("ExecuteCode() error = %v", err)
	}

	if result.Success {
		t.Error("Expected execution to fail due to timeout")
	}

	if !strings.Contains(result.Error, "timeout") {
		t.Errorf("Expected timeout error, got: %s", result.Error)
	}
}
