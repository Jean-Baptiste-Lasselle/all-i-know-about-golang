data "pokus_projects" "example2" {
  depends_on = [
    pokus_project.godzilla_project,
    pokus_project.mothra_project,
    pokus_project.gidhora_project,
  ]
}
