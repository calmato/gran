provider "google" {
  project     = var.project_id
  region      = "asia-northeast1"
  credentials = file(var.credential_file)
}

module "this" {
  source = "./../../modules"

  location = "asia-northeast1-a"

  #################################################
  # GKE Cluster
  #################################################
  gke_cluster_name        = "gran_cluster"
  gke_cluster_description = "gran application cluster for staging"

  gke_cluster_min_master_version = "1.14.10-gke.17"

  #################################################
  # GKE Node
  #################################################
  gke_node_name_prefix = "gran_node"

  gke_node_count        = 1
  gke_node_machine_type = "f1-micro"

  tags = {
    service = "gran"
    env     = "stg"
  }
}
