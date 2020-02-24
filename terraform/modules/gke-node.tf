#################################################
# GKE Node
#################################################
resource "google_container_node_pool" "this" {
  name_prefix = var.gke_node_name_prefix

  cluster    = google_container_cluster.this.name
  location   = var.location

  node_count = var.gke_node_count

  node_config {
    preemptible  = true
    machine_type = var.gke_node_machine_type

    oauth_scopes = [
      "https://www.googleapis.com/auth/cloud-platform"
    ]
  }
}
