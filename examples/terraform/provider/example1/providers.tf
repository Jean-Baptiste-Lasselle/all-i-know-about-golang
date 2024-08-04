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
