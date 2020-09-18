resource "google_compute_address" "back_end" {
    depends_on = [google_project_service.enabled_services]

    name = "back-end"
}

resource "google_compute_address" "front_end" {
    depends_on = [google_project_service.enabled_services]

    name = "front-end"
}