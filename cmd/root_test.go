package cmd

import (
	"bytes"
	"strings"
	"testing"
)

func TestRootCommand(t *testing.T) {
	buffer := new(bytes.Buffer)
	rootCmd.SetOut(buffer)
	rootCmd.SetArgs([]string{})

	t.Cleanup(func() {
		rootCmd.SetOut(nil)
		rootCmd.SetArgs(nil)
	})

	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("execute root command: %v", err)
	}

	if got := buffer.String(); got != "Hello from your-cli\n" {
		t.Fatalf("unexpected output: %q", got)
	}
}

func TestCompletionCommand(t *testing.T) {
	buffer := new(bytes.Buffer)
	rootCmd.SetOut(buffer)
	rootCmd.SetArgs([]string{"completion", "bash"})

	t.Cleanup(func() {
		rootCmd.SetOut(nil)
		rootCmd.SetArgs(nil)
	})

	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("execute completion command: %v", err)
	}

	if got := buffer.String(); !strings.Contains(got, "bash completion for your-cli") {
		t.Fatalf("unexpected completion output: %q", got)
	}
}
