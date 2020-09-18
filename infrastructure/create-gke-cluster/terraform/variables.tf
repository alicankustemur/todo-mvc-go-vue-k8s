variable "project" {
    description = "The project id"
}

variable "region" {
    description = "region"
}

variable "ip_cidr_range" {
    description = "The CIDR range of cluster"
}

variable "initial_node_count" {
    description = "The number of nodes to create in this cluster's default node pool."
}

variable "node_count" {
    description = "number of gke nodes"
}

variable "machine_type" {
    description = "the type of virtual machines"
}

variable "services" {
    description = "Terraform needs to be authorized to communicate with following the Google Cloud APIs"
    type        = list(string)
    
    default = [
        "container.googleapis.com",
        "compute.googleapis.com",
        "servicenetworking.googleapis.com",
        "cloudresourcemanager.googleapis.com",
    ]
}

# It role only can change existing resources
# and pull GCR images
variable "gke_service_account_roles" {
    description = "List of roles to be granted to the GKE cluster."
    type        = list(string)
    
    default = [
        "roles/editor",
        "roles/storage.objectViewer"
    ]
}