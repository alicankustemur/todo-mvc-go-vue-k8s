resource "null_resource" "get_credentials" {
    depends_on = [google_container_cluster.primary]

    triggers = {
        always_run = "${timestamp()}"
    }

    provisioner "local-exec" {
        command = <<EOT
            gcloud auth activate-service-account --key-file="$GOOGLE_APPLICATION_CREDENTIALS" --project=${var.project}
            gcloud config set core/project ${var.project} 
            gcloud config set compute/region ${var.region}
            gcloud config set compute/region ${var.region}
            gcloud container clusters get-credentials ${google_container_cluster.primary.name} --region ${var.region}
        EOT
    }
}

resource "local_file" "service_account_key_json" {
    content  = base64decode(google_service_account_key.gke.private_key)
    filename = "${path.module}/${google_service_account.gke.account_id}.json"
}

resource "null_resource" "enable_pull_gcr_images" {
    depends_on = [null_resource.get_credentials]
    
    triggers = {
        always_run = "${timestamp()}"
    }

    provisioner "local-exec" {
        command = <<EOT
            kubectl create secret docker-registry ${google_service_account.gke.account_id} \
                --docker-server=https://eu.gcr.io \
                --docker-username=_json_key \
                --docker-email=user@example.com \
                --docker-password="$(cat ${path.module}/${google_service_account.gke.account_id}.json )"
            
            kubectl patch serviceaccount default \
                -p "{\"imagePullSecrets\": [{\"name\": \"${google_service_account.gke.account_id}\"}]}"
        EOT
    }
}