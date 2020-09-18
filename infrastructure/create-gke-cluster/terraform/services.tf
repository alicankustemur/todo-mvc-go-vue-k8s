resource "google_project_service" "enabled_services" {
    project = var.project
    count   = length(var.services)
    service = element(var.services, count.index)
    
    disable_on_destroy  = false
}