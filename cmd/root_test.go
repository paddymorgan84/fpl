package cmd

import (
	"testing"
)

func TestBuildRootCommand(t *testing.T) {
	cmd := BuildRootCommand()

	var expectedShort = `
		________    _______      ___
		|   ____|   |   _  \     |  |
		|  |__      |  |_)  |    |  |
		|   __|     |   ___/     |  |
		|  |        |  |         |  -----.
		|__|        |__|         |_______|

		A CLI tool for retrieving FPL data`
	if cmd.Short != expectedShort {
		t.Fatalf("expected: %v; got %v", expectedShort, cmd.Short)
	}

	var expectedUse = "fpl"
	if cmd.Use != expectedUse {
		t.Fatalf("expected: %v; got %v", expectedUse, cmd.Use)
	}
}
