provider "google" {
  project     = var.project_id
  region      = "asia-northeast1"
  credentials = file("account.json")
}

module "this" {
  source = "./../../modules"

  location = "asia-northeast1-a"

  gke_cluster_name        = "gran_master"
  gke_cluster_description = "Gran GKE Master"
}
