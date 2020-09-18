# Create the GKE cluster
resource "google_container_cluster" "primary" {
  name     = "gke"
  location = var.region

  remove_default_node_pool = true
  initial_node_count       = var.initial_node_count

  min_master_version = data.google_container_engine_versions.gke_version.latest_master_version

  network     = google_compute_network.vpc.name
  subnetwork  = google_compute_subnetwork.subnet.name

  master_auth {
    username = ""
    password = ""

    client_certificate_config {
      issue_client_certificate = false
    }
  }
}

# Separately Managed Node Pool
resource "google_container_node_pool" "primary_nodes" {
  name       = "${google_container_cluster.primary.name}-node-pool"
  location   = var.region
  cluster    = google_container_cluster.primary.name
  node_count = var.node_count

  node_config {
    service_account = google_service_account.gke.email

    oauth_scopes = [
      "https://www.googleapis.com/auth/logging.write",
      "https://www.googleapis.com/auth/monitoring",
    ]

    labels = {
      env = var.project
    }

    machine_type = var.machine_type
    tags         = ["gke-node", "${var.project}-gke"]
    metadata = {
      disable-legacy-endpoints = "true"
    }
  }
}