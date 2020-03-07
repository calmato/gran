provider "google" {
  project = var.project_id
  region  = "asia-northeast1"
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
  gke_node_configs = [
    {
      name  = "gran-node"
      count        = 1
      preemptible  = false
      machine_type = "f1-micro"
    },
    {
      name  = "gran-preemptible-node"
      count        = 1
      preemptible  = true
      machine_type = "n1-standard-1"
    },
  ]
}
