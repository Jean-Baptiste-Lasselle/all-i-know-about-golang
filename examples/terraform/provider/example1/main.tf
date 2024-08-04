resource "pokus_project" "godzilla_project" {
  name                 = "godzilla"
  description          = "a first example project to test creating a project with OpenTOFU the terraformation king. It also has been updated using the OpenTOFU King."
  git_service_provider = "giteaJapan"
  git_ssh_uri          = "git@github.com:3forges/godzilla.git"
}

resource "pokus_project" "mothra_project" {
  name                 = "mothra"
  description          = "a second example project to test creating a project with OpenTOFU the terraformation king. It also has been updated using the OpenTOFU King."
  git_service_provider = "giteaJapan"
  git_ssh_uri          = "git@github.com:3forges/mothra.git"

}

resource "pokus_project" "gidhora_project" {
  name                 = "gidhora"
  description          = "a third example project to test creating a project with OpenTOFU the terraformation king"
  git_service_provider = "giteaJapan"
  git_ssh_uri          = "git@github.com:3forges/gidhora.git"
}