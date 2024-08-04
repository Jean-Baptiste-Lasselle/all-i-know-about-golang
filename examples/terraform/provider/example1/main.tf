resource "pokus_project" "example_project" {
  /*
  items = [
    {
      description          = "a third example project to test updating a project with curl"
      git_service_provider = "giteaJapan"
      git_ssh_uri          = "git@github.com:3forges/gidhora.git"
      name                 = "gidhora"
    }
  ]
  */
  // items = []
  description          = "a third example project to test updating a project with curl"
  git_service_provider = "giteaJapan"
  git_ssh_uri          = "git@github.com:3forges/gidhora.git"
  name                 = "gidhora"
}
