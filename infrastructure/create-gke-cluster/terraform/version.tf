terraform {
  required_version = ">= 0.12"
}

# Get avaiable GKE version in the region
data "google_container_engine_versions" "gke_version" {
  depends_on = [google_project_service.enabled_services]

  project  = var.project
  location = var.region
}