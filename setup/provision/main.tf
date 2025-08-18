### Terraform Configuration

terraform {
  required_providers {
    hcloud = {
      source  = "hetznercloud/hcloud"
      version = "~> 1.45"
    }
  }
}

variable "hcloud_token" {
  sensitive = true
  type      = string
}

variable "ssh_pubkey_path" {
  sensitive = true
  type      = string
}

variable "root_vm_config" {
  description = <<-EOT
	Configuration for Root VM
  EOT
  type = object({
    name        = string
    image       = string
    server_type = string
    location    = string
  })

  default = {
    name        = "root-vm"
    image       = "ubuntu-22.04"
    server_type = "cx22"
    location    = "nbg1"
  }

}

provider "hcloud" {
  token = var.hcloud_token
}


### Resource Definitions

resource "hcloud_ssh_key" "this" {
  name       = "my-ssh-key"
  public_key = file(var.ssh_pubkey_path)
}

resource "hcloud_server" "this" {
  name        = var.root_vm_config.name
  image       = var.root_vm_config.image
  server_type = var.root_vm_config.server_type
  location    = var.root_vmn_config.location

  ssh_keys = [hcloud_ssh_key.this.id]
}
