terraform {
  required_providers {
    pokus = {
      source = "pokus-io.io/terraform/pokus"
    }
  }
}

provider "pokus" {
    // host = "http://api.pesto.io:3000"
    // username = "education"
    // password = "sdsddg"
  host     = "http://api.pesto.io:3000"
  username = "education"
  password = "test123"
}

data "pokus_projects" "example" {}


output "pokus_projects" {
  description = "access key as root to the minio storage backend of the csi driver."
  value       = data.pokus_projects.example
}