data "pokus_projects" "example2" {
  depends_on = [pokus_project.example_project]
}
