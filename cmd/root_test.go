package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func newTestRootCommand(buffer *bytes.Buffer, args ...string) *cobra.Command {
	command := NewRootCommand("test")
	command.SetOut(buffer)
	command.SetErr(buffer)
	command.SetArgs(args)
	return command
}

func TestRootCommand(t *testing.T) {
	buffer := new(bytes.Buffer)
	command := newTestRootCommand(buffer)

	if err := command.Execute(); err != nil {
		t.Fatalf("execute root command: %v", err)
	}

	if got := buffer.String(); got != "Hello from your-cli\n" {
		t.Fatalf("unexpected output: %q", got)
	}
}

func TestCompletionCommand(t *testing.T) {
	buffer := new(bytes.Buffer)
	command := NewRootCommand("test")

	if err := command.GenBashCompletionV2(buffer, true); err != nil {
		t.Fatalf("generate bash completion: %v", err)
	}

	if got := buffer.String(); !strings.Contains(got, "# bash completion V2 for your-cli") {
		t.Fatalf("unexpected completion output: %q", got)
	}
}

func TestVersionCommand(t *testing.T) {
	buffer := new(bytes.Buffer)
	command := newTestRootCommand(buffer, "version")

	if err := command.Execute(); err != nil {
		t.Fatalf("execute version command: %v", err)
	}

	if got := buffer.String(); got != "your-cli test\n" {
		t.Fatalf("unexpected output: %q", got)
	}
}
