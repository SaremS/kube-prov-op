package test

import (
	"testing"
	"os"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformHetznerRootVmProvision(t *testing.T) {
	t.Parallel()

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options {
		TerraformDir: "../provision"
		Vars: map[string]interface{}{}
	})
}
