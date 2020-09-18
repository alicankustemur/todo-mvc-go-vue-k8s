output "cluster_name" {
    value = google_container_cluster.primary.name
}

output "backend_external_ip" {
    value = google_compute_address.back_end.address
}

output "frontend_external_ip" {
    value = google_compute_address.front_end.address
}