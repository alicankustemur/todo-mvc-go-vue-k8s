# Create the dedicated GKE service account for the application cluster
resource "google_service_account" "gke" {
    account_id   = "${var.project}-gke"
    display_name = "${var.project}-gke"
    project      = var.project
}

resource "google_service_account_key" "gke" {
    service_account_id = google_service_account.gke.name
    public_key_type    = "TYPE_X509_PEM_FILE"
}

# Bind permissions to GKE service account
resource "google_project_iam_member" "gke" {
    count   = length(var.gke_service_account_roles)
    project = var.project
    role    = element(var.gke_service_account_roles, count.index)
    member  = format("serviceAccount:%s", google_service_account.gke.email)
}