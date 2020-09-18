provider "google" {
  project     = var.project
  region      = var.region 
  version     = "~> 3.23.0"
}

provider "google-beta" {
  project     = var.project
  region      = var.region 
  version     = "~> 3.23.0"
}

provider "null" {
  version = "~> 2.1.2"
}

provider "local" {
  version = "~> 1.4.0"
}