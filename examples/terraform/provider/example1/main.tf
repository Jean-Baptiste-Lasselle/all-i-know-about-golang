resource "pokus_project" "godzilla_project" {
  name                 = "godzillaDemoCedric"
  description          = "DEMO CEDRIC a first example project to test creating a project with OpenTOFU the terraformation king. It also has been updated using the OpenTOFU King. A third test updating a project, to test if [stringplanmodifier.UseStateForUnknown()] works."
  git_service_provider = "giteaJapan"
  git_ssh_uri          = "git@github.com:3forges/godzillaDemoCedric.git"
}

resource "pokus_project" "mothra_project" {
  name                 = "mothraDemoCedric"
  description          = "DEMO CEDRIC a second example project to test creating a project with OpenTOFU the terraformation king. It also has been updated using the OpenTOFU King. A fourth test updating a project, to test if [stringplanmodifier.UseStateForUnknown()] works."
  git_service_provider = "giteaJapan"
  git_ssh_uri          = "git@github.com:3forges/mothra.git"

}

resource "pokus_project" "gidhora_project" {
  name                 = "gidhoraDemoCedric"
  description          = "DEMO CEDRIC a third example project to test creating a project with OpenTOFU the terraformation king"
  git_service_provider = "giteaJapan"
  git_ssh_uri          = "git@github.com:3forges/gidhora.git"
}