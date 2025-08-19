package test

import (
	"fmt"
	"net"
	"os"
	"testing"
	"io/ioutil"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestTerraformHetznerRootVmProvision(t *testing.T) {
	t.Parallel()
	
	tempDir, err := ioutil.TempDir("", "ssh-key-")
	if err != nil {
		panic(fmt.Sprintf("Failed to create temp directory: %v", err))
	}
	defer os.RemoveAll(tempDir)

	privKey, pubKey, err := generateSshKeys()
	if err != nil {
		panic(fmt.Sprintf("Failed to generate keypair: %v", err))
	}

	publicKeyPath, err :=storeSshKeys(privKey, pubKey, tempDir)
	if err != nil {
		panic(fmt.Sprintf("Failed to store keypair: %v", err))
	}

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../../setup/provision/hetzner",
		Vars: map[string]interface{}{
			"hcloud_token":    os.Getenv("HCLOUD_TOKEN"),
			"ssh_pubkey_path": publicKeyPath,
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
