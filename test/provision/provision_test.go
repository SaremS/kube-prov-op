package test

import (
	"net"
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestTerraformHetznerRootVmProvision(t *testing.T) {
	t.Parallel()

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../../setup/provision/hetzner",
		Vars: map[string]interface{}{
			"hcloud_token":    os.Getenv("HCLOUD_TOKEN"),
			"ssh_pubkey_path": os.Getenv("SSH_PUBKEY_PATH"),
		},
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	ip4Address := terraform.Output(t, terraformOptions, "vm_ip4_address")
	ip4Parsed := net.ParseIP(ip4Address)
	if ip4Parsed == nil || ip4Parsed.To4() == nil {
		t.Errorf("Expected %q to be a valid IPv4 address, but it was not.", ip4Address)
	}
}
