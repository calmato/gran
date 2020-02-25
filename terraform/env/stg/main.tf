provider "google" {
  project     = var.project_id
  region      = "asia-northeast1"
}

module "this" {
  source = "./../../modules"

  location = "asia-northeast1-a"

  #################################################
  # GKE Cluster
  #################################################
  gke_cluster_name        = "gran-cluster"
  gke_cluster_description = "gran application cluster for staging"

  gke_cluster_min_master_version = "1.14.10-gke.17"

  #################################################
  # GKE Node
  #################################################
  gke_node_name_prefix = "gran-node"

  gke_node_count        = 1
  gke_node_machine_type = "f1-micro"
}
