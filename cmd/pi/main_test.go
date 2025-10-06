package main_test

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/initialcapacity/pi/internal/assert"
)

func TestIntegration(t *testing.T) {
	prepareBuildDirectory(t)
	build(t, "pi")

	testCtx, cancelCtx := context.WithCancel(t.Context())
	defer cancelCtx()

	output := runCommand(t, testCtx, "./build/pi", "-d", "1", "-n", "2")

	assert.ContainsSubstring(t, output, "Running for 1")
	assert.ContainsSubstring(t, output, "with 2 workers")
	assert.ContainsSubstring(t, output, "π ≈ 3.14")
}

func prepareBuildDirectory(t *testing.T) {
	err := os.RemoveAll("../../build")
	assert.NoError(t, err)
	err = os.MkdirAll("./build", os.ModePerm)
	assert.NoError(t, err)
}

func build(t *testing.T, name string) {
	err := exec.Command("go", "build", "-o", fmt.Sprintf("./build/%s", name), fmt.Sprintf("../../cmd/%s", name)).Run()
	assert.NoError(t, err)
}

func runCommand(t *testing.T, ctx context.Context, command string, arguments ...string) string {
	output, err := exec.CommandContext(ctx, command, arguments...).Output()
	assert.NoError(t, err)
	return string(output)
}
